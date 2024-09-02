package repository

import "loan_apps/datasource/domain/entity"

type PartnerRepository interface {
	Save(partner *entity.Partner) ( error)
	List() ([]*entity.Partner, error)
	FindById(id string) (*entity.Partner, error)
	Update(partner *entity.Partner) (error)
	Delete(id string) error
}