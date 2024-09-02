package unittesting

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/schema"
	"loan_apps/internal/services"
	mocktesting "loan_apps/test/mock_testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoanServiceCreate(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)

	request := schema.LoanReq{
		CustomerID: "1",
		TenorID: "1",
	}

	IDTenure :="1"
	expectedTenor := &entity.Tenor{
		ID: "1",
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockTenureRepo.On("FindById", IDTenure).Return(expectedTenor, nil)

	mockLoanRepo.On("Save", mock.Anything).Return(nil)
	err := mockService.Create(&request)
	assert.NoError(t, err)
	mockLoanRepo.AssertExpectations(t)
}

func TestLoanServiceUpdate(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)

	request := schema.LoanReq{
		CustomerID: "1",
		TenorID: "1",
	}

	IDLoan := "1"
	expectedLoan := &entity.Loan{
		ID: "1",
		CustomerID: "1",
		Amount: 1000000,
		InstalmentMonth: 1,
	}

	IDTenure :="1"
	expectedTenor := &entity.Tenor{
		ID: "1",
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockLoanRepo.On("FindById", IDLoan).Return(expectedLoan, nil)

	mockTenureRepo.On("FindById", IDTenure).Return(expectedTenor, nil)

	mockLoanRepo.On("Update", mock.Anything).Return(nil)
	err := mockService.Update(expectedLoan.ID,&request)

	assert.NoError(t, err)
	mockLoanRepo.AssertExpectations(t)
}

func TestLoanServiceDelete(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)

	IDLoan := "1"

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockLoanRepo.On("Delete", IDLoan).Return(nil)

	err := mockService.Delete(IDLoan)

	assert.NoError(t, err)
	mockLoanRepo.AssertExpectations(t)
}

func TestLoanServiceFindById(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)
	
	IDLoan := "1"

	expectedLoan := &entity.Loan{
		ID: "1",
		CustomerID: "1",
		Amount: 1000000,
		InstalmentMonth: 1,
	}

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockLoanRepo.On("FindById", IDLoan).Return(expectedLoan, nil)

	loan, err := mockService.FindByID(IDLoan)

	assert.NoError(t, err)
	assert.Equal(t, expectedLoan, loan)
	mockLoanRepo.AssertExpectations(t)
}

func TestLoanServiceList(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)

	expectedLoan := []*entity.Loan{
		{
			ID: "1",
			CustomerID: "1",
			Amount: 1000000,
			InstalmentMonth: 1,
		},
	}

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockLoanRepo.On("List").Return(expectedLoan, nil)

	loan, err := mockService.List()
	
	assert.NoError(t, err)
	assert.Equal(t, expectedLoan, loan)
	mockLoanRepo.AssertExpectations(t)
}

func TestLoanServiceFindByCustomerID(t *testing.T) {
	mockLoanRepo := new(mocktesting.MockLoanRepository)
	mockTenureRepo := new(mocktesting.MockTenorRepository)

	IDCustomer := "1"
	expectedLoan := []*entity.Loan{
		{
			ID: "1",
			CustomerID: "1",
			Amount: 1000000,
			InstalmentMonth: 1,
		},
	}

	mockService := services.NewtLoanService(mockLoanRepo, mockTenureRepo)

	mockLoanRepo.On("FindByCustomerID", IDCustomer).Return(expectedLoan, nil)

	loan, err := mockService.FindByCustomerID(IDCustomer)
	
	assert.NoError(t, err)
	assert.Equal(t, expectedLoan, loan)
	mockLoanRepo.AssertExpectations(t)
}