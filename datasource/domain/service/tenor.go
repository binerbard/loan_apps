package service

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
)

type TenorService interface {
	Create(credential *schema.TenorReq) error
	Delete(id string) error
	Update(id string, credential *schema.TenorReq) error
	FindByID(id string) (*entity.Tenor, error)
	List() ([]*entity.Tenor, error)
	FindByCustomerID(id string) ([]*entity.Tenor, error)
}