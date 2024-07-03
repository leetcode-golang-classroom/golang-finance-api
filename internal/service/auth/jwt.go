package auth

import (
	"context"
	"time"

	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == jwtWare.ErrJWTMissingOrMalformed.Error() {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
}
func AuthMiddle(secret string) fiber.Handler {
	// setup jwt route
	return jwtWare.New(
		jwtWare.Config{
			SigningKey:   jwtWare.SigningKey{Key: []byte(secret)},
			ErrorHandler: errorHandler,
		},
	)
}

func (h *Handler) CreateToken(ctx context.Context, username string) (string, error) {
	// Create jwt claim
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().UTC().Add(time.Hour * 1).Unix(),
		"iat": time.Now().UTC().Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.jwtSigneSecret))
}
