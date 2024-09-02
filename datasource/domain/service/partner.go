package service

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
)

type PartnerService interface {
	FindByID(id string) (*entity.Partner, error)
	List() ([]*entity.Partner, error)
	Create(credential *schema.PartnerReq) error
	Update(id string, credential *schema.PartnerReq) error
	Delete(id string) error
}