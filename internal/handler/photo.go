package handler

import (
	"strconv"

	"github.com/SantiiL/unsplash-image-viewer-backend/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
)

type PhotoHandler struct {
	service unsplash.UnsplashService
}

func NewPhotoHandler(service unsplash.UnsplashService) *PhotoHandler {
	return &PhotoHandler{service: service}
}

func (h *PhotoHandler) GetPhotos(c *fiber.Ctx) error {
	pageParam := c.Query("page", "1")
	perPageParam := c.Query("per_page", "12")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.Status(400).SendString("Page parameter is not a valid integer")
	}

	perPage, err := strconv.Atoi(perPageParam)
	if err != nil {
		return c.Status(400).SendString("Per_page parameter is not a valid integer")
	}

	photos, err := h.service.GetPhotos(page, perPage)
	if err != nil {
		return c.Status(500).SendString("Failed to get photos from Unsplash")
	}

	return c.JSON(photos)
}
