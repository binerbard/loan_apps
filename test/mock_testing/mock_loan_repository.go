package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)


type MockLoanRepository struct {
	mock.Mock
}


// Save(loan *entity.Loan) ( error)
// Update(loan *entity.Loan) (error)
// Delete(id string) error
// FindById(id string) (*entity.Loan, error)
// List() ([]*entity.Loan, error)
// FindByCustomerID(customer_id string) ([]*entity.Loan, error)

func (r *MockLoanRepository) Save(loan *entity.Loan, balance *entity.Balance) ( error) {
	return r.Called(loan).Error(0)
}

func (r *MockLoanRepository) Update(loan *entity.Loan) (error) {
	return r.Called(loan).Error(0)
}

func (r *MockLoanRepository) Delete(id string) (error) {
	return r.Called(id).Error(0)
}

func (r *MockLoanRepository) FindById(id string) (*entity.Loan, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.Loan), args.Error(1)
}

func (r *MockLoanRepository) List() ([]*entity.Loan, error) {
	return r.Called().Get(0).([]*entity.Loan), r.Called().Error(1)
}

func (r *MockLoanRepository) FindByCustomerID(customer_id string) ([]*entity.Loan, error) {
	args := r.Called(customer_id)
	return args.Get(0).([]*entity.Loan), args.Error(1)
}