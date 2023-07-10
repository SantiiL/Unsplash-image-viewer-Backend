package handler

import (
	"strconv"

	"github.com/SantiiL/unsplash-image-viewer-backend/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
)

type SearchHandler struct {
	service unsplash.UnsplashService
}

func NewSearchHandler(service unsplash.UnsplashService) *SearchHandler {
	return &SearchHandler{service: service}
}

func (h *SearchHandler) SearchPhotos(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return c.Status(400).SendString("Query parameter is required")
	}

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

	searchResponse, err := h.service.SearchPhotos(query, page, perPage)
	if err != nil {
		return c.Status(500).SendString("Failed to search photos from Unsplash")
	}

	return c.JSON(searchResponse)
}
