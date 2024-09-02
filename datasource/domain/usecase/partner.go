package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
)


type PartnerUsecase interface {
	service.PartnerService
}


type partnerUsecase struct {
	partnerRepository repository.PartnerRepository
	partnerService service.PartnerService
}

func NewPartnerUsecase(partnerRepository repository.PartnerRepository, partnerService service.PartnerService) PartnerUsecase {
	return &partnerUsecase{
		partnerRepository: partnerRepository,
		partnerService: partnerService,
	}
}

func (u *partnerUsecase) Create(credential *schema.PartnerReq) error {
	return u.partnerService.Create(credential)
}


func (u *partnerUsecase) Update(id string, credential *schema.PartnerReq) error {
	return u.partnerService.Update(id, credential)
}


func (u *partnerUsecase) Delete(id string) error {
	return u.partnerService.Delete(id)
}

func (u *partnerUsecase) FindByID(id string) (*entity.Partner, error) {
	return u.partnerService.FindByID(id)
}

func (u *partnerUsecase) List() ([]*entity.Partner, error) {
	return u.partnerService.List()
}