package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)


type TenorRepositoryImpl struct {
	DB *gorm.DB
}

func NewTenorRepository(db *gorm.DB) repository.TenorRepository {
	return &TenorRepositoryImpl{DB: db}
}



func (r *TenorRepositoryImpl) Save(tenor *entity.Tenor) ( error){
	if err := r.DB.Create(tenor).Error ; err != nil {
		return err
	}
	return nil
}
func (r *TenorRepositoryImpl) Update(tenor *entity.Tenor) (error){
	if err := r.DB.Save(tenor).Error; err != nil {
		return err
	}
	return nil
}
func (r *TenorRepositoryImpl) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Tenor{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *TenorRepositoryImpl) FindById(id string) (*entity.Tenor, error) {
	Tenor := &entity.Tenor{}
	if err := r.DB.Where("id = ?", id).First(Tenor).Error; err != nil {
		return nil, err
	}
	return Tenor, nil
}

func (r *TenorRepositoryImpl) List() ([]*entity.Tenor, error){
	var Tenors []*entity.Tenor
	if err := r.DB.Find(&Tenors).Error; err != nil {
		return nil, err
	}
	return Tenors, nil
}