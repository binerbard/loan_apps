package repository

import "loan_apps/datasource/domain/entity"



type BalanceRepository interface {
	Save(balance *entity.Balance) ( error)
	Update(balance *entity.Balance) (error)
	Delete(id string) error
	FindById(id string) (*entity.Balance, error)
	FindByCustomerID(id string) (*entity.Balance, error)
	List() ([]*entity.Balance, error)
}