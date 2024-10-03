package routes

import (
	"multiaura/internal/controllers"
	"multiaura/internal/repositories"
	"multiaura/internal/services"
	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app *fiber.App) {
	repository := repositories.NewUserRepository(mongoDB)
	service := services.NewUserService(repository)
	controller := controllers.NewUserController(service)

	userGroup := app.Group("/user")

    userGroup.Post("/register", controller.Register)
    userGroup.Post("/login", controller.Login)
}