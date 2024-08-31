package repository

import "loan_apps/datasource/domain/entity"



type TransactionRepository interface {
	Save(transaction *entity.Transaction) ( error)
	Update(transaction *entity.Transaction) (error)
	Delete(id string) error
	FindById(id string) (*entity.Transaction, error)
	List() ([]*entity.Transaction, error)
}