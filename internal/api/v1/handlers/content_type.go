package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunquan03/cms_api/internal/models"
)

func (h *Handler) CreateContentType(c *fiber.Ctx) error {
	var contentType models.ContentType
	if err := c.BodyParser(&contentType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	err := h.service.CreateContentType(&contentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.SuccessResp{
		Status:  true,
		Message: "success",
	})
}
