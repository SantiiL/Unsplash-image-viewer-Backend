package handler

import (
	"strconv"

	"github.com/SantiiL/unsplash-image-viewer-backend/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
)

func GetPhotos(c *fiber.Ctx) error {
	pageParam := c.Query("page", "1")         // Default to page 1 if not provided
	perPageParam := c.Query("per_page", "12") // Default to 10 items per page if not provided

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.Status(400).SendString("Page parameter is not a valid integer")
	}

	perPage, err := strconv.Atoi(perPageParam)
	if err != nil {
		return c.Status(400).SendString("Per_page parameter is not a valid integer")
	}

	photos, err := unsplash.GetPhotos(page, perPage)
	if err != nil {
		// Log the error and return a 500 status code with a helpful message
		return c.Status(500).SendString("Failed to get photos from Unsplash")
	}

	return c.JSON(photos)
}
