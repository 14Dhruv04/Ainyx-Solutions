package main

import (
	"log"

	"ainyx/config"
	"ainyx/db/sqlc"
	"ainyx/internal/handler"
	"ainyx/internal/logger"
	"ainyx/internal/repository"
	"ainyx/internal/routes"
	"ainyx/internal/service"
	"ainyx/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := logger.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Log.Sync()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)

	repo := repository.NewUserRepository(queries)

	userService := service.NewUserService(repo)

	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Use(middleware.RequestLogger)

	routes.RegisterRoutes(app, userHandler)

	log.Fatal(app.Listen(":3000"))
}
