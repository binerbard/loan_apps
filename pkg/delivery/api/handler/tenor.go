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

func (h *TenorHandler) List(ctx *fiber.Ctx) error {
	tenors, err := h.tenorUsecase.List()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": tenors,
	})
}

func (h *TenorHandler) Find(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	tenor, err := h.tenorUsecase.FindByID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": tenor,
	})
}

func (h *TenorHandler) Delete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	if err := h.tenorUsecase.Delete(ID); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success deleted tenor",
	})
}

func (h *TenorHandler) Update(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

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

	if err := h.tenorUsecase.Update(ID, &req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success updated tenor",
	})
}

func (h *TenorHandler) FindByCustomerID(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	tenors, err := h.tenorUsecase.FindByCustomerID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": tenors,
	})
}