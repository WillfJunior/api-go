package main

import (
	"github.com/WillfJunior/api-go/controllers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Enable Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/clients", controllers.GetClients)
	e.GET("/clients/:id", controllers.GetClient)
	e.POST("/clients", controllers.CreateClient)
	e.PUT("/clients/:id", controllers.UpdateClient)
	e.DELETE("/clients/:id", controllers.DeleteClient)

	// Start the server
	e.Start(":8080")
}
