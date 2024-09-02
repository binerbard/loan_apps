package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}


func (r *MockTransactionRepository) Save(transaction *entity.Transaction, dealer *entity.Dealer) ( error) {
	args := r.Called(transaction, dealer)
	return args.Error(0)
}

func (r *MockTransactionRepository) Update(transaction *entity.Transaction) (error) {
	args := r.Called(transaction)
	return args.Error(0)
}


func (r *MockTransactionRepository) Delete(id string) error {
	return r.Called(id).Error(0)
}

func (r *MockTransactionRepository) FindById(id string) (*entity.Transaction, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (r *MockTransactionRepository) List() ([]*entity.Transaction, error) {
	args := r.Called()
	return args.Get(0).([]*entity.Transaction), args.Error(1)
	
}