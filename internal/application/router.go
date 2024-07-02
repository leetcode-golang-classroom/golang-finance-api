package application

func (app *App) SetupRoutes() {
	api := app.fiberApp.Group("/api")
	// setup handler
	handler := NewHandler(app.db)
	// health check
	api.Get("/", handler.HealthCheck)
}
