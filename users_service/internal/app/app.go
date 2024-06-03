package app

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"users_service/internal/handlers"
)

func Run() {
	server := fiber.New()

	handlers.Register(server)

	log.Fatal(server.Listen(":3000"))
}
