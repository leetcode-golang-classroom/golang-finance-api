package application

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util/response"
)

type Handler struct {
	db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	var result int
	err := h.db.Get(&result, "select 1")
	if err != nil {
		return errors.New("database unavailable")
	}
	return response.Ok(c, fiber.Map{
		"database": "available",
	})
}
