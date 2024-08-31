package service

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
)


type CustomerService interface {
	Create(credential *schema.CustomerReq) error
	Delete(id string) error
	Update(id string, credential *schema.CustomerReq) error
	FindByID(id string) (*entity.Customer, error)
	List() ([]*entity.Customer, error)
}