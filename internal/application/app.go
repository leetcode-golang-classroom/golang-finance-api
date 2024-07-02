package application

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/config"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/db"
	"github.com/leetcode-golang-classroom/golang-finance-api/internal/util"
)

type App struct {
	db       *sqlx.DB
	fiberApp *fiber.App
	config   *config.Config
}

func New(config *config.Config) *App {
	dbInstance, err := db.Connect(config.DbURL)
	if err != nil {
		util.FailOnError(err, "failed to connect")
	}
	fiberApp := fiber.New()
	app := &App{
		fiberApp: fiberApp,
		db:       dbInstance,
		config:   config,
	}
	app.loadRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	log.Printf("Starting server on %d\n", app.config.Port)
	addr := fmt.Sprintf(":%d", app.config.Port)
	// error buffer channel
	errCh := make(chan error, 1)
	// setup app listener
	go func() {
		err := app.fiberApp.Listen(addr)
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	// teardown connection
	defer app.Stop()
	// listen for errCh andd context donee event
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done(): // for cancel
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return app.fiberApp.ShutdownWithContext(timeout)
	}
}
func (app *App) Stop() {
	if err := app.db.Close(); err != nil {
		log.Println("failed to close db connection", err)
	}
}
