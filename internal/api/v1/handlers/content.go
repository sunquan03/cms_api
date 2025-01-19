package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sunquan03/cms_api/internal/models"
	"strconv"
)

func (h *Handler) CreateContent(c *fiber.Ctx) error {
	contentType := c.Params("content_type")
	content := make(map[string]interface{})
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	id, err := h.service.CreateContent(contentType, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResp{
		Status:  true,
		Message: "success",
		Result:  id,
	})
}

func (h *Handler) GetContentById(c *fiber.Ctx) error {
	contentType := c.Params("content_type")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: "invalid id format",
		})
	}

	jsonData, err := h.service.GetContentById(contentType, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(jsonData)
}

func (h *Handler) UpdateContent(c *fiber.Ctx) error {
	contentType := c.Params("content_type")
	content := make(map[string]interface{})
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: "invalid id format",
		})
	}

	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = h.service.UpdateContent(contentType, id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResp{
		Status:  true,
		Message: "success",
	})
}

func (h *Handler) DeleteContent(c *fiber.Ctx) error {
	contentType := c.Params("content_type")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: "invalid id format",
		})
	}

	err = h.service.DeleteContent(contentType, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResp{
		Status:  true,
		Message: "success",
	})
}

func (h *Handler) SearchContentByQuery(c *fiber.Ctx) error {
	contentType := c.Params("content_type")
	searchQuery := c.Query("query")
	if len(searchQuery) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: "empty search query",
		})
	}

	data, err := h.service.SearchContentByQuery(context.TODO(), contentType, searchQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
