package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/types"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util"
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
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	password, err := h.passwordHdr.HashPassword(input.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid password")
	}
	user, err := h.userStore.Create(ctx.Context(), input.Username, password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "register error")
	}
	token, err := h.CreateToken(ctx.Context(), user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "register error")
	}
	return util.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	input := types.AuthRequest{}
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	user, err := h.userStore.GetUserByName(ctx.Context(), input.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}
	if !h.passwordHdr.CheckPassword(input.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "login failded")
	}
	token, err := h.CreateToken(ctx.Context(), user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "register error")
	}
	return util.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) Me(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*jwt.Token)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "token not found")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "claim failed")
	}
	username, ok := claims["sub"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "claim structure error")
	}
	currentUser, err := h.userStore.GetUserByName(ctx.Context(), username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}
	token, err := h.CreateToken(ctx.Context(), currentUser.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}
	return util.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})

}
