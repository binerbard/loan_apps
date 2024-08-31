package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
	"time"

	"github.com/google/uuid"
)


type TenorServiceImpl struct {
	tenorData repository.TenorRepository
}

func NewtTenorService(tenorRepository repository.TenorRepository) service.TenorService {
	return &TenorServiceImpl{
		tenorData: tenorRepository,
	}
}

func (s *TenorServiceImpl) Create(credential *schema.TenorReq) error {
	tenor := entity.Tenor{
		ID: uuid.New().String(),
		CustomerID: credential.CustomerID,
		Amount: credential.Amount,
		InstalmentMonth: credential.InstalmentMonth,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.tenorData.Save(&tenor)

}
func (s *TenorServiceImpl) Delete(id string) error {
	return s.tenorData.Delete(id)
}

func (s *TenorServiceImpl) Update(id string, credential *schema.TenorReq) error {
	tenor, err := s.tenorData.FindById(id)
	if err != nil {
		return err
	}

	tenor.Amount = credential.Amount
	tenor.InstalmentMonth = credential.InstalmentMonth
	tenor.UpdatedAt = time.Now()

	return s.tenorData.Update(tenor)
}

func (s *TenorServiceImpl) FindByID(id string) (*entity.Tenor, error) {
	return s.tenorData.FindById(id)
}

func (s *TenorServiceImpl) List() ([]*entity.Tenor, error){
	return s.tenorData.List()
}