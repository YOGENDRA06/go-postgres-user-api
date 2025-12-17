package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"Go_Backend_Development_Task/config"
	"Go_Backend_Development_Task/db/sqlc"
	"Go_Backend_Development_Task/internal/handler"
	"Go_Backend_Development_Task/internal/logger"
	"Go_Backend_Development_Task/internal/middleware"
	"Go_Backend_Development_Task/internal/repository"
	"Go_Backend_Development_Task/internal/routes"
	"Go_Backend_Development_Task/internal/service"
)


func main() {
	// ---------------- Logger ----------------
	logger.Init()
	defer logger.Log.Sync()

	// ---------------- DB ----------------
	db, err := sql.Open("postgres", config.DBUrl())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)

	// ---------------- Layers ----------------
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	// ---------------- Fiber App ----------------
	app := fiber.New(fiber.Config{
		AppName: "Go Backend Development Task",
	})

	// ---------------- Middleware ----------------
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	// ---------------- Routes ----------------
	routes.Register(app, userHandler)

	logger.Log.Info("server started", zap.String("port", "8080"))
	log.Fatal(app.Listen("0.0.0.0:8080"))
}
