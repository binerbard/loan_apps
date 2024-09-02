package repository

import "loan_apps/datasource/domain/entity"



type LoanRepository interface {
	Save(loan *entity.Loan, balance *entity.Balance) ( error)
	Update(loan *entity.Loan) (error)
	Delete(id string) error
	FindById(id string) (*entity.Loan, error)
	List() ([]*entity.Loan, error)
	FindByCustomerID(customer_id string) ([]*entity.Loan, error)
}