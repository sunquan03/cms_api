package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunquan03/cms_api/internal/api/v1/handlers"
)

func Routes(app *fiber.App, handler *handlers.Handler) {
	v1 := app.Group("/api/v1")

	v1.Post("/content_types", handler.CreateContenttype)
}
