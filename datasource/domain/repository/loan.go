package repository

import "loan_apps/datasource/domain/entity"



type LoanRepository interface {
	Save(loan *entity.Loan) ( error)
	Update(loan *entity.Loan) (error)
	Delete(id string) error
	FindById(id string) (*entity.Loan, error)
	List() ([]*entity.Loan, error)
}