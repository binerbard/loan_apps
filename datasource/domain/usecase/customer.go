package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
)



type CustomerUsecase interface {
	service.CustomerService
}

type customerUsecase struct {
	customerRepository repository.CustomerRepository
	customerService service.CustomerService
}

func NewUCustomerUsecase(r repository.CustomerRepository, s service.CustomerService) CustomerUsecase {
	return &customerUsecase{
		customerRepository: r,
		customerService: s,
	}
}

func(u *customerUsecase) Create(credential *schema.CustomerReq) error{
	return u.customerService.Create(credential)
}

func(u *customerUsecase) Delete(id string) error{
	return u.customerService.Delete(id)
}

func(u *customerUsecase) Update(id string, credential *schema.CustomerReq) error {
	return u.customerService.Update(id, credential)
}

func(u *customerUsecase) FindByID(id string) (*entity.Customer, error){
	return u.customerService.FindByID(id)
}

func (u *customerUsecase) List() ([]*entity.Customer, error){
	return u.customerService.List()
}