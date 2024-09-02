package service

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
)

type LoanService interface {
	Create(credential *schema.LoanReq) error
	Delete(id string) error
	Update(id string, credential *schema.LoanReq) error
	FindByID(id string) (*entity.Loan, error)
	List() ([]*entity.Loan, error)
	FindByCustomerID(id string) ([]*entity.Loan, error)
}