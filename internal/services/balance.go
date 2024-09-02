package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
)

type BalanceServiceImpl struct {
	balanceRepo repository.BalanceRepository
}

func NewBalanceService(balanceRepo repository.BalanceRepository) *BalanceServiceImpl {
	return &BalanceServiceImpl{balanceRepo: balanceRepo}
}

func (s *BalanceServiceImpl) List() ([]*entity.Balance, error) {
	return s.balanceRepo.List()
}

func (s *BalanceServiceImpl) FindByID(id string) (*entity.Balance, error) {
	return s.balanceRepo.FindById(id)
}

func (s *BalanceServiceImpl) FindByCustomerID(id string) (*entity.Balance, error) {
	return s.balanceRepo.FindByCustomerID(id)
}