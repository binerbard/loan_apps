package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)

type PartnerRepositoryImpl struct {
	DB *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) repository.PartnerRepository {
	return &PartnerRepositoryImpl{DB: db}
}

func (p *PartnerRepositoryImpl) Save(partner *entity.Partner) ( error) {
	if err := p.DB.Create(partner).Error; err != nil {
		return err
	}
	return nil
}

func (p *PartnerRepositoryImpl) Delete(id string) error {
	var partner entity.Partner
	if err := p.DB.Where("id = ?", id).Delete(&partner).Error; err != nil {
		return err
	}
	return nil
}

func (p *PartnerRepositoryImpl) List() ([]*entity.Partner, error) {
	var partners []*entity.Partner
	if err := p.DB.Find(&partners).Error; err != nil {
		return nil, err
	}
	return partners, nil
}

func (p *PartnerRepositoryImpl) FindById(id string) (*entity.Partner, error) {
	var partner entity.Partner
	if err := p.DB.Where("id = ?", id).Find(&partner).Error; err != nil {
		return nil, err
	}
	return &partner, nil
}


func (p *PartnerRepositoryImpl) Update(partner *entity.Partner) (error) {
	if err := p.DB.Save(partner).Error; err != nil {
		return err
	}
	return nil
}