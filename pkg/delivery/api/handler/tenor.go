package handler

import (
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)

type TenorHandler struct {
	tenorUsecase usecase.TenorUsecase
}

func NewTenorHandler(uc usecase.TenorUsecase) *TenorHandler {
	return &TenorHandler{
		tenorUsecase: uc,
	}
}


func (h *TenorHandler) Store(ctx *fiber.Ctx) error {	
	var req schema.TenorReq

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorTenorReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.tenorUsecase.Create(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
	})
}
