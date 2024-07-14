package response

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewApiError(err error, code int, message string, data interface{}) *ApiError {
	log.Printf("error: %v\n", err)
	return &ApiError{
		Code:    code,
		Message: message,
		Data:    &data,
	}
}
func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	e, ok := err.(*ApiError)
	if !ok {
		ef, ok := err.(*fiber.Error)
		if !ok {
			e = NewApiError(err, fiber.StatusInternalServerError, e.Error(), nil)
		} else {
			e = NewApiError(err, ef.Code, ef.Error(), nil)
		}
	}
	return ctx.Status(e.Code).JSON(ApiResponse{
		Success: false,
		Error:   e,
		Meta:    GetMetaData(ctx),
	})
}

func GetMetaData(ctx *fiber.Ctx) ApiMetaData {
	return ApiMetaData{
		Timestamp: time.Now(),
		Path:      ctx.Path(),
		Method:    ctx.Method(),
	}
}
func Response(ctx *fiber.Ctx, code int, data any) error {
	return ctx.Status(code).JSON(ApiResponse{
		Success: true,
		Data:    data,
		Meta:    GetMetaData(ctx),
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}

func ErrorBadRequest(err error) error {
	return NewApiError(
		err,
		fiber.StatusBadRequest,
		err.Error(),
		nil,
	)
}

func ErrorUnauthorized(err error, message string) error {
	return NewApiError(
		err,
		fiber.StatusUnauthorized,
		message,
		nil,
	)
}
