package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHandler interface {
	CheckPassword(password, hash string) bool
	HashPassword(password string) (string, error)
}

type Handler struct{}

func (h *Handler) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *Handler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
