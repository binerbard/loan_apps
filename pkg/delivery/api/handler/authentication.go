package handler

import (
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler struct {
	authUseCase usecase.UserUsecase
}

func NewAuthenticationHandler(uc usecase.UserUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{
		authUseCase: uc,
	}
}

func (h *AuthenticationHandler) SignIn(ctx *fiber.Ctx) error {
	var req schema.SignInReq

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidateSignInReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.authUseCase.(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
	})
}


func (h *AuthenticationHandler) SignUp(ctx *fiber.Ctx) error {
	var req schema.SignUpReq

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidateSignUpReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.authUseCase.CreateUser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}
}