package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
	"time"

	"github.com/google/uuid"
)


type CustomerServiceImpl struct {
	customerData repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) service.CustomerService {
	return &CustomerServiceImpl{
		customerData: customerRepository,
	}
}

func (s *CustomerServiceImpl) Create(credential *schema.CustomerReq) error {
	customer := entity.Customer{
		ID: uuid.New().String(),
		NIK: credential.NIK,
		FullName: credential.FullName,
		LegalName: credential.LegalName,
		BirthDay: credential.BirthDay,
		BirthPlace: credential.BirthPlace,
		Salary: credential.Salary,
		ImgID: credential.ImgID,
		ImgSelfie: credential.ImgSelfie,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.customerData.Save(&customer)

}
func (s *CustomerServiceImpl) Delete(id string) error {
	return s.customerData.Delete(id)
}

func (s *CustomerServiceImpl) Update(id string, credential *schema.CustomerReq) error {
	customer, err := s.customerData.FindById(id)
	if err != nil {
		return err
	}

	customer.NIK = credential.NIK
	customer.FullName = credential.FullName
	customer.LegalName = credential.LegalName
	customer.BirthDay = credential.BirthDay
	customer.BirthPlace = credential.BirthPlace
	customer.Salary = credential.Salary
	if credential.ImgSelfie != "" {
		customer.ImgSelfie = credential.ImgSelfie
	}
	if credential.ImgID != "" {
		customer.ImgID = credential.ImgID
	}

	customer.UpdatedAt = time.Now()

	return s.customerData.Update(customer)
}

func (s *CustomerServiceImpl) FindByID(id string) (*entity.Customer, error) {
	return s.customerData.FindById(id)
}

func (s *CustomerServiceImpl) List() ([]*entity.Customer, error){
	return s.customerData.List()
}