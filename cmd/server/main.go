package main

import (
	"log"

	"github.com/SantiiL/unsplash-image-viewer-backend/internal/handler"
	"github.com/SantiiL/unsplash-image-viewer-backend/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create an instance of the Unsplash service
	unsplashService := unsplash.NewUnsplashService()

	// Create an instance of your handlers and pass in the service
	photoHandler := handler.NewPhotoHandler(unsplashService)
	searchHandler := handler.NewSearchHandler(unsplashService)

	app := fiber.New()

	// Configurar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Use the new handler instances for your routes
	app.Get("/photos", photoHandler.GetPhotos)
	app.Get("/search/photos", searchHandler.SearchPhotos)

	app.Listen(":3010")
}
