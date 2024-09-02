package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
	"time"

	"github.com/google/uuid"
)

type PartnerServiceImpl struct {
	partnerData repository.PartnerRepository
}

func NewPartnerService(r repository.PartnerRepository) service.PartnerService {
	return &PartnerServiceImpl{partnerData: r}
}

func (s *PartnerServiceImpl) Create(credential *schema.PartnerReq) error {
	partner := &entity.Partner{
		Name: credential.Name,
		ID: uuid.NewString(),
		Contact: credential.Contact,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.partnerData.Save(partner)
}

func (s *PartnerServiceImpl) List() ([]*entity.Partner, error) {
	return s.partnerData.List()
}

func (s *PartnerServiceImpl) FindByID(id string) (*entity.Partner, error) {
	return s.partnerData.FindById(id)
}

func (s *PartnerServiceImpl) Update(id string,credential *schema.PartnerReq) (error) { 
	partner, err := s.partnerData.FindById(id)
	if err != nil {
		return err
	}
	partner.Name = credential.Name
	partner.Contact = credential.Contact
	partner.UpdatedAt = time.Now()
	return s.partnerData.Update(partner)
}


func (s *PartnerServiceImpl) Delete(id string) error {
	return s.partnerData.Delete(id)
}

