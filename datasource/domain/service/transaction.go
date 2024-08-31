package service

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
)

type TransactionService interface {
	Create(credential *schema.TransactionReq) error
	Delete(id string) error
	Update(id string, credential *schema.TransactionReq) error
	FindByID(id string) (*entity.Transaction, error)
	List() ([]*entity.Transaction, error)
}