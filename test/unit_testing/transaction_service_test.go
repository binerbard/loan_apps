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

func TestTransactionServiceList(t *testing.T) {

	mockTransactionRepo := new(mocktesting.MockTransactionRepository)

	data :=[] *entity.Transaction{
		{
		ID: "1",
		CustomerID: "1",
		ContractNumber: "1",
		OTR: 1,
		AdminFee: 1,
		InstallmentAmount: 1,
		InterestAmount: 1,
		AssetName: "1",
	},
	}

	mockService := services.NewtTransactionService(mockTransactionRepo)
	mockTransactionRepo.On("List").Return(data, nil)

	transaction, err := mockService.List()

	assert.NoError(t, err)

	mockTransactionRepo.AssertExpectations(t)
	assert.Equal(t, data, transaction)
}

func TestTransactionServiceFindByID(t *testing.T) {
	mockTransactionRepo := new(mocktesting.MockTransactionRepository)
	ID := "1"
	expectedTransaction := &entity.Transaction{
		ID: "1",
		CustomerID: "1",
		ContractNumber: "1",
		OTR: 1,
		AdminFee: 1,
		InstallmentAmount: 1,
		InterestAmount: 1,
		AssetName: "1",
	}
	mockService := services.NewtTransactionService(mockTransactionRepo)
	mockTransactionRepo.On("FindById", ID).Return(expectedTransaction, nil)
	response, err := mockService.FindByID(ID)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	mockTransactionRepo.AssertExpectations(t)

	assert.Equal(t, "1", response.ID)
	assert.Equal(t, "1", response.CustomerID)
	assert.Equal(t, "1", response.ContractNumber)
	assert.Equal(t, 1, response.OTR)
	assert.Equal(t, 1, response.AdminFee)
	assert.Equal(t, 1, response.InstallmentAmount)
	assert.Equal(t, 1, response.InterestAmount)

}

func TestTransactionCreate(t *testing.T) {
	mockTransactionRepo := new(mocktesting.MockTransactionRepository)

	request := &schema.TransactionReq{
		CustomerID:      "1",
		ContractNumber:  "1",
		PartnerID:       "1",
		OTR:             1,
		AdminFee:        1,
		InstallmentAmount: 1,
		InterestAmount:  1,
		AssetName:       "1",
	}

	// Create transaction and dealer mock values
	transactionMatcher := mock.MatchedBy(func(tx *entity.Transaction) bool {
		return tx.CustomerID == request.CustomerID &&
			tx.ContractNumber == request.ContractNumber &&
			tx.OTR == request.OTR &&
			tx.AdminFee == request.AdminFee &&
			tx.InstallmentAmount == request.InstallmentAmount &&
			tx.InterestAmount == request.InterestAmount &&
			tx.AssetName == request.AssetName
	})

	dealerMatcher := mock.MatchedBy(func(dl *entity.Dealer) bool {
		return dl.PartnerID == request.PartnerID
	})

	mockService := services.NewtTransactionService(mockTransactionRepo)

	// Expect Save to be called with the matched arguments
	mockTransactionRepo.On("Save", transactionMatcher, dealerMatcher).Return(nil)

	// Call the service method
	mockService.Create(request)

	// Assert that expectations were met
	mockTransactionRepo.AssertExpectations(t)
}
