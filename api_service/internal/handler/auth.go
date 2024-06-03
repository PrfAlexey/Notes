package handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"api_service/internal/entity"
)

type authUsecase interface {
	SignUp(ctx context.Context, req *entity.SignUpRequest) error
}

type auth struct {
	usecase authUsecase
}

// SignUp Регистрация пользователя
// @Summary   Регистрация пользователя
// @Tags      auth
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Accept    json
// @Produce   json
// @Param     signUpRequest  body  entity.SignUpRequest  true  "Запрос на регистрацию пользователя"
// @Success   200  "OK"
// @Failure   400
// @Failure   401
// @Failure   403
// @Failure   404
// @Failure   409
// @Failure   500
// @Router    /sign-up [post]
func (h auth) SignUp(c *fiber.Ctx) error {
	req := entity.SignUpRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := h.usecase.SignUp(c.Context(), &req); err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}
