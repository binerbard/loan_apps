package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/service"
)

type BalanceUsecase interface {
	service.BalanceService
}

type balanceUsecase struct {
	balanceRepository repository.BalanceRepository
	balanceService	service.BalanceService
}

func NewBalanceUsecase(balanceRepository repository.BalanceRepository, balanceService service.BalanceService) BalanceUsecase {
	return &balanceUsecase{
		balanceRepository: balanceRepository,
		balanceService:	balanceService,
	}
}


func (u *balanceUsecase) FindByID(id string) (*entity.Balance, error) {
	return u.balanceService.FindByID(id)
}

func (u *balanceUsecase) List() ([]*entity.Balance, error) {
	return u.balanceService.List()
}

func (u *balanceUsecase) FindByCustomerID(id string) (*entity.Balance, error) {
	return u.balanceService.FindByCustomerID(id)
}