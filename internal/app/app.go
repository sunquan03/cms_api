package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/sunquan03/cms_api/internal/api/v1"
	"github.com/sunquan03/cms_api/internal/api/v1/handlers"
	"github.com/sunquan03/cms_api/internal/cache"
	"github.com/sunquan03/cms_api/internal/models"
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"github.com/sunquan03/cms_api/internal/repository/postgres"
	"github.com/sunquan03/cms_api/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const idleTimeout = 5 * time.Second

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	_port := os.Getenv("app_port")
	_host := os.Getenv("app_host")

	log.Printf("Dynamic Content Management System - verison %s\n", os.Getenv("app_version"))
	_app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			return c.Status(code).JSON(models.ErrorResp{
				Status:  false,
				Message: err.Error(),
			})
		},
		DisableStartupMessage: true,
	})

	_app.Use(recover.New(), cors.New(), logger.New())

	_db := postgres.NewDB()
	_esClient, _esTransport := elastic.NewESClient()
	_redisClient := cache.NewRedisClient()

	_postgresRepo := postgres.NewPostgresLayer(_db)
	_elasticRepo := elastic.NewElasticLayer(_esClient, _esTransport)
	_redisRepo := cache.NewRedisCache(_redisClient)

	syncChan := make(chan *models.ContentSync, 100)

	_service := service.NewService(_elasticRepo, _postgresRepo, _redisRepo, syncChan)

	syncWorkerPool := service.NewSWPool(10, syncChan, _elasticRepo)

	_handler := handlers.NewHandler(_service)

	v1.Routes(_app, _handler)

	go syncWorkerPool.Run()

	go func() {
		if err := _app.Listen(fmt.Sprintf("%s:%s", _host, _port)); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c
	_ = _app.Shutdown()

	close(syncChan)
	_db.Close()
	_elasticRepo.Close()

}
