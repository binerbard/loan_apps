package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockBalanceRepository struct {
	mock.Mock
}


func (m *MockBalanceRepository) Save(balance *entity.Balance) ( error) {
	return m.Called(balance).Error(0)
}


func (m *MockBalanceRepository) Update(balance *entity.Balance) (error) {
	return m.Called(balance).Error(0)
}

func (m *MockBalanceRepository) Delete(id string) error {
	return m.Called(id).Error(0)
}

func (m *MockBalanceRepository) FindById(id string) (*entity.Balance, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *MockBalanceRepository) FindByCustomerID(id string) (*entity.Balance, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *MockBalanceRepository) List() ([]*entity.Balance, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Balance), args.Error(1)
}