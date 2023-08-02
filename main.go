package main

import (
	handlers "dasso/open-api/Handlers"
	"log"

	"github.com/labstack/echo/v4"
	//"github.com/open-api/handlers"
)

func main() {
	e := echo.New()

	// Initialize the API handler
	apiHandler := handlers.NewAPIHandler()

	// Register the generated OpenAPI handlers
	e.GET("/hello", apiHandler.GetHello)

	// Start the server
	log.Println("Server listening at :8080")
	e.Start(":8080")
}
