package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)

// MockCustomerRepo adalah mock dari CustomerRepo
type MockCustomerRepo struct {
    mock.Mock
}

// Implementasikan metode-metode pada MockCustomerRepo
func (m *MockCustomerRepo) Save(customer *entity.Customer) error {
    args := m.Called(customer)
    return args.Error(0)
}

func (m *MockCustomerRepo) Update(customer *entity.Customer) error {
    args := m.Called(customer)
    return args.Error(0)
}

func (m *MockCustomerRepo) Delete(id string) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *MockCustomerRepo) FindById(id string) (*entity.Customer, error) {
    args := m.Called(id)
    return args.Get(0).(*entity.Customer), args.Error(1)
}

func (m *MockCustomerRepo) List() ([]*entity.Customer, error) {
    args := m.Called()
    return args.Get(0).([]*entity.Customer), args.Error(1)
}
