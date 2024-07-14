package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/types"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util/response"
	"github.com/leetcode-golang-classroom/golang-finance-api/pkg/password"
)

type Handler struct {
	userStore      types.UserStore
	passwordHdr    password.PasswordHandler
	jwtSigneSecret string
}

func NewAuthRouter(userStore types.UserStore, passwordHdr password.PasswordHandler,
	jwtSignedSecret string,
) *Handler {
	return &Handler{userStore, passwordHdr, jwtSignedSecret}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	router.Post("/login", h.Login)
	router.Post("/register", h.Register)
	router.Get("/me", AuthMiddle(h.jwtSigneSecret), h.Me)
}
func (h *Handler) Register(ctx *fiber.Ctx) error {
	input := types.AuthRequest{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	password, err := h.passwordHdr.HashPassword(input.Password)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	user, err := h.userStore.Create(ctx.Context(), input.Username, password)
	if err != nil {
		return response.ErrorUnauthorized(err, "register error")
	}
	token, err := h.CreateToken(ctx.Context(), user.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "register error")
	}
	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	input := types.AuthRequest{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	user, err := h.userStore.GetUserByName(ctx.Context(), input.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "login error")
	}
	if !h.passwordHdr.CheckPassword(input.Password, user.Password) {
		return response.ErrorUnauthorized(err, "login failded")
	}
	token, err := h.CreateToken(ctx.Context(), user.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "login error")
	}
	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) Me(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*jwt.Token)
	if !ok {
		return response.ErrorBadRequest(fiber.NewError(fiber.StatusBadRequest, "token not found"))
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return response.ErrorBadRequest(fiber.NewError(fiber.StatusBadRequest, "claim failed"))
	}
	username, ok := claims["sub"].(string)
	if !ok {
		return response.ErrorBadRequest(fiber.NewError(fiber.StatusBadRequest, "claim structure error"))
	}
	currentUser, err := h.userStore.GetUserByName(ctx.Context(), username)
	if err != nil {
		return response.ErrorUnauthorized(err, "invalid token")
	}
	token, err := h.CreateToken(ctx.Context(), currentUser.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "invalid token")
	}
	return response.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
