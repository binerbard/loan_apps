package handler

import (
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)

type PartnerHandler struct {
	PartnerUsecase usecase.PartnerUsecase
}

func NewPartnerHandler(uc usecase.PartnerUsecase) *PartnerHandler {
	return &PartnerHandler{
		PartnerUsecase: uc,
	}
}

func (h *PartnerHandler) Store(ctx *fiber.Ctx) error {

	var req schema.PartnerReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorPartnerReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.PartnerUsecase.Create(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success created partner",
	})

}

func (h *PartnerHandler) Find(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	partner, err := h.PartnerUsecase.FindByID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": partner,
	})
}

func (h *PartnerHandler) Update(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	var req schema.PartnerReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorPartnerReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.PartnerUsecase.Update(ID, &req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success updated partner",
	})
}

func (h *PartnerHandler) Delete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	if err := h.PartnerUsecase.Delete(ID); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success deleted partner",
	})
}

func (h *PartnerHandler) List(ctx *fiber.Ctx) error {
	partners, err := h.PartnerUsecase.List()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": partners,
	})
}