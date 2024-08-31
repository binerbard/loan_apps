package repository

import "loan_apps/datasource/domain/entity"


type TenorRepository interface {
	Save(tenor *entity.Tenor) ( error)
	Update(tenor *entity.Tenor) (error)
	Delete(id string) error
	FindById(id string) (*entity.Tenor, error)
	List() ([]*entity.Tenor, error)
}