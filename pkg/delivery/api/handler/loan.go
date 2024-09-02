package handler

import (
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	loanUsecase usecase.LoanUsecase
}

func NewLoanHandler(uc usecase.LoanUsecase) *LoanHandler {
	return &LoanHandler{
		loanUsecase: uc,
	}
}


func (h *LoanHandler) Store(ctx *fiber.Ctx) error {	
	var req schema.LoanReq

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorLoanReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.loanUsecase.Create(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
	})
}

func (h *LoanHandler) List(ctx *fiber.Ctx) error {
	loans, err := h.loanUsecase.List()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": loans,
	})
}

func (h *LoanHandler) Find(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	loan, err := h.loanUsecase.FindByID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": loan,
	})
}

func (h *LoanHandler) Delete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	if err := h.loanUsecase.Delete(ID); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success deleted loan",
	})
}

func (h *LoanHandler) Update(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	var req schema.LoanReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := schema.ValidatorLoanReq(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.loanUsecase.Update(ID, &req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success updated loan",
	})
}


func (h *LoanHandler) FindByCustomerID(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	loans, err := h.loanUsecase.FindByCustomerID(ID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": loans,
	})
}