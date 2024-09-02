package service

import (
	"loan_apps/datasource/domain/entity"
)

type BalanceService interface {
	FindByID(id string) (*entity.Balance, error)
	FindByCustomerID(id string) (*entity.Balance, error)
	List() ([]*entity.Balance, error)
}