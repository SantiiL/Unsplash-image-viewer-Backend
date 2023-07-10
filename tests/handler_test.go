package handler_test

import (
	"net/http"
	"testing"

	"github.com/SantiiL/unsplash-image-viewer-backend/internal/handler"
	"github.com/SantiiL/unsplash-image-viewer-backend/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock implementation for UnsplashService
type UnsplashServiceMock struct {
	mock.Mock
}

func (m *UnsplashServiceMock) GetPhotos(page int, perPage int) ([]unsplash.Photo, error) {
	args := m.Called(page, perPage)
	return args.Get(0).([]unsplash.Photo), args.Error(1)
}

func (m *UnsplashServiceMock) SearchPhotos(query string, page int, perPage int) (unsplash.SearchResponse, error) {
	args := m.Called(query, page, perPage)
	return args.Get(0).(unsplash.SearchResponse), args.Error(1)
}

func TestGetPhotos(t *testing.T) {
	// Mock UnsplashService to simulate GetPhotos behaviour
	unsplashMock := new(UnsplashServiceMock)
	unsplashMock.On("GetPhotos", 1, 12).Return([]unsplash.Photo{}, nil)

	// Initialize Fiber router
	app := fiber.New()
	handler := handler.NewPhotoHandler(unsplashMock)
	app.Get("/photos", handler.GetPhotos)

	// Create a GET request to /photos
	req, _ := http.NewRequest("GET", "/photos", nil)
	resp, _ := app.Test(req)

	require.Equal(t, 200, resp.StatusCode, "Should return a HTTP 200 status")

	unsplashMock.AssertExpectations(t)
}

func TestSearchPhotos(t *testing.T) {
	// Mock UnsplashService to simulate SearchPhotos behaviour
	unsplashMock := new(UnsplashServiceMock)
	unsplashMock.On("SearchPhotos", "testQuery", 1, 12).Return(unsplash.SearchResponse{}, nil)

	// Initialize Fiber router
	app := fiber.New()
	handler := handler.NewSearchHandler(unsplashMock)
	app.Get("/search", handler.SearchPhotos)

	// Create a GET request to /search
	req, _ := http.NewRequest("GET", "/search?query=testQuery", nil)
	resp, _ := app.Test(req)

	require.Equal(t, 200, resp.StatusCode, "Should return a HTTP 200 status")

	unsplashMock.AssertExpectations(t)
}
