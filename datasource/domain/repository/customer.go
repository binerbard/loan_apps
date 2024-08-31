package repository

import "loan_apps/datasource/domain/entity"



type CustomerRepository interface {
	Save(customer *entity.Customer) ( error)
	Update(customer *entity.Customer) (error)
	Delete(id string) error
	FindById(id string) (*entity.Customer, error)
	List() ([]*entity.Customer, error)
}