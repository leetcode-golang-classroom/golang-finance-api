package application

import (
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/service/auth"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/service/user"
	"github.com/leetcode-golang-classroom/golang-finance-api/pkg/password"
)

func (app *App) SetupRoutes() {
	api := app.fiberApp.Group("/api")
	// setup handler
	handler := NewHandler(app.db)
	// health check
	api.Get("/", handler.HealthCheck)
	// setup userStore
	userStore := user.NewUserStore(app.db)
	// setup passwordHdr
	passwordHdr := &password.Handler{}
	// setup auth handler
	authHandler := auth.NewAuthRouter(userStore, passwordHdr, app.config.JWTSignSecret)
	// setup auth api
	authHandler.SetupRoutes(api)
}
