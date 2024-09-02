package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
)



type TenorUsecase interface {
	service.TenorService
}

type tenorUsecase struct {
	tenorRepository repository.TenorRepository
	tenorService service.TenorService
}

func NewUTenorUsecase(r repository.TenorRepository, s service.TenorService) TenorUsecase {
	return &tenorUsecase{
		tenorRepository: r,
		tenorService: s,
	}
}

func(u *tenorUsecase) Create(credential *schema.TenorReq) error{
	return u.tenorService.Create(credential)
}

func(u *tenorUsecase) Delete(id string) error{
	return u.tenorService.Delete(id)
}

func(u *tenorUsecase) Update(id string, credential *schema.TenorReq) error {
	return u.tenorService.Update(id, credential)
}

func(u *tenorUsecase) FindByID(id string) (*entity.Tenor, error){
	return u.tenorService.FindByID(id)
}

func (u *tenorUsecase) List() ([]*entity.Tenor, error){
	return u.tenorService.List()
}

func (u *tenorUsecase) FindByCustomerID(id string) ([]*entity.Tenor, error){
	return u.tenorService.FindByCustomerID(id)
}