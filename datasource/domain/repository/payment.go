package repository

import "loan_apps/datasource/domain/entity"



type PaymentRepository interface {
	Save(payment *entity.Payment) ( error)
	Update(payment *entity.Payment) (error)
	Delete(id string) error
	FindById(id string) (*entity.Payment, error)
	List() ([]*entity.Payment, error)
}