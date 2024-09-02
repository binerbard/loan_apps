package unittesting

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/internal/services"
	mocktesting "loan_apps/test/mock_testing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestBalanceServiceFind(t *testing.T) {

	mockBalance := new(mocktesting.MockBalanceRepository)

	ID :="1"
	expectedBalance := &entity.Balance{
		ID: ID,
		CustomerID: "1",
		Balance: 0,
	}

	
	mockService := services.NewBalanceService(mockBalance)
	mockBalance.On("FindById", mock.Anything).Return(expectedBalance, nil)
	
	balance, err := mockService.FindByID(ID)
	
	assert.NoError(t, err)
	
	mockBalance.AssertExpectations(t)
	assert.Equal(t, expectedBalance, balance)
	
}

func TestBalanceServiceList(t *testing.T) {
	mockBalance := new(mocktesting.MockBalanceRepository)
	
	data :=[] *entity.Balance{
		{
		ID: "1",
		CustomerID: "1",
		Balance: 0,
	},
	}

	mockService := services.NewBalanceService(mockBalance)
	mockBalance.On("List").Return(data, nil)
	
	balance, err := mockService.List()
	
	assert.NoError(t, err)
	
	mockBalance.AssertExpectations(t)
	assert.Equal(t, data, balance)

}

func TestBalanceServiveFindByCustomerID(t *testing.T) {

	mockBalance := new(mocktesting.MockBalanceRepository)
	ID := "1"
	expectedBalance := &entity.Balance{
		ID: ID,
		CustomerID: "1",
		Balance: 0,
	}
	mockService := services.NewBalanceService(mockBalance)
	mockBalance.On("FindByCustomerID", mock.Anything).Return(expectedBalance, nil)
	
	balance, err := mockService.FindByCustomerID(ID)
	
	assert.NoError(t, err)
	
	mockBalance.AssertExpectations(t)
	assert.Equal(t, expectedBalance, balance)
	
}