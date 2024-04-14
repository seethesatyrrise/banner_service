package app

import (
	"bannerService/internal/cache"
	"bannerService/internal/config"
	"bannerService/internal/database"
	"bannerService/internal/deletion"
	"bannerService/internal/handler"
	"bannerService/internal/repo"
	"bannerService/internal/service"
	"bannerService/internal/utils"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type App struct {
	db            *database.DB
	cache         *cache.Cache
	deletionQueue *deletion.DeletionQueue
	router        *echo.Echo
	cfg           *config.Config
	services      *service.Service
}

func New() (app *App, err error) {

	app = &App{}

	utils.Logger.Info("config initializing")

	app.cfg, err = config.New()
	if err != nil {
		return nil, errors.Wrap(err, "reading config err")
	}

	app.db, err = database.New(&app.cfg.DB)
	if err != nil {
		return nil, errors.Wrap(err, "database connection err")
	}

	app.cache = cache.New(&app.cfg.Cache)

	app.deletionQueue = deletion.CreateQueue()

	log.Info("database connected")

	app.router = echo.New()

	repos := repo.New(app.db.DB)
	app.services = service.New(repos, app.cache, app.deletionQueue)
	handlers := handler.New(app.services, &app.cfg.Tokens)

	handlers.Route(app.router)

	return app, err
}

func (app *App) Run() error {

	log.Info("server starting")

	return app.router.Start(":" + app.cfg.HTTP.Port)
}

func (app *App) Shutdown(ctx context.Context) error {
	if err := app.services.Close(); err != nil {
		return err
	}
	return app.router.Shutdown(ctx)
}
