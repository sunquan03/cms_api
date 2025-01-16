package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunquan03/cms_api/internal/api/v1/handlers"
)

func Routes(app *fiber.App, handler *handlers.Handler) {
	v1 := app.Group("/api/v1")

	v1.Post("/content_types", handler.CreateContentType)

	content := v1.Group("/content/:content_type")

	content.Post("/", handler.CreateContent)
	content.Put("/:id", handler.UpdateContent)
	// TODO: content.Delete("/:id", handler.DeleteContent)
	content.Get("/:id", handler.GetContentById)

}
