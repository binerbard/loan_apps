package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockTenorRepository struct {
	mock.Mock
}

func (r *MockTenorRepository) Save(tenor *entity.Tenor) ( error) {
	return r.Called(tenor).Error(0)
}

func (r *MockTenorRepository) AvailableInstalment(customer_id string, instalment_month int) (*entity.Tenor, error) {
	args := r.Called(customer_id, instalment_month)
	return args.Get(0).(*entity.Tenor), args.Error(1)
}

func (r *MockTenorRepository) Update(tenor *entity.Tenor) (error) {
	return r.Called(tenor).Error(0)
}

func (r *MockTenorRepository) Delete(id string) (error) {
	return r.Called(id).Error(0)
}

func (r *MockTenorRepository) FindById(id string) (*entity.Tenor, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.Tenor), args.Error(1)
}

func (r *MockTenorRepository) List() ([]*entity.Tenor, error) {
	return r.Called().Get(0).([]*entity.Tenor), r.Called().Error(1)
}

func (r *MockTenorRepository) FindByCustomerID(id string) ([]*entity.Tenor, error) {
	args := r.Called(id)
	return args.Get(0).([]*entity.Tenor), args.Error(1)
}