package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.4
//go:generate swag init -o ../../docs/swagger/v1 -g router.go --parseInternal --parseDependency  --instanceName v1

//go:generate go install github.com/go-swagger/go-swagger/cmd/swagger@latest
//go:generate swagger generate client -f ../../docs/swagger/v1/v1_swagger.json -t ../../pkg/apiclient/v1 -m entity

const (
	usersURL = "/api/users"
	userURL  = "/api/users/:uuid"
)

// @title        users service API
// @version      1.0
// @description  Сервис юзеров.
// @BasePath     /v1
func Register(server *fiber.App) {

	// Создание юзера
	server.Post(usersURL, func(ctx *fiber.Ctx) error {
		return nil
	})
	// Получение юзера
	server.Get(usersURL, func(ctx *fiber.Ctx) error { return nil })

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	server.Get("/info", GetAppInfo)

}

// GetAppInfo возвращает информацию о приложении
// @Description  Получает информацию о приложении
// @Summary      Получить информацию о приложении
// @Tags         info
// @Produce      json
// @Success      200
// @Failure      500
// @ID           info
// @Router       /info [get]
func GetAppInfo(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
