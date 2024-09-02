package handler

import (
	"loan_apps/datasource/domain/usecase"
	"loan_apps/scripts/utils"

	"github.com/gofiber/fiber/v2"
)


type BalanceHandler struct {
	balanceUsecase usecase.BalanceUsecase
}

func NewBalanceHandler(uc usecase.BalanceUsecase) *BalanceHandler {
	return &BalanceHandler{
		balanceUsecase: uc,
	}
}

func (u *BalanceHandler) List(ctx *fiber.Ctx) ( error) {
	balances, err := u.balanceUsecase.List()
	if err != nil {
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": balances,
	})

}

func (u *BalanceHandler) Find(ctx *fiber.Ctx) ( error) {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	balance, err := u.balanceUsecase.FindByID(ID)
	if err != nil {
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": balance,
	})
}

func (u *BalanceHandler) FindByCustomerID(ctx *fiber.Ctx) ( error) {
	ID := ctx.Params("id")
	if ID == "" {
		return utils.BadRequest(ctx, fiber.Map{
			"message": "id is empty",
		})
	}

	balance, err := u.balanceUsecase.FindByCustomerID(ID)
	if err != nil {
		return utils.BadRequest(ctx, fiber.Map{
			"message": err.Error(),
		})
	}

	return utils.Success(ctx, fiber.Map{
		"message": "success",
		"data": balance,
	})
}