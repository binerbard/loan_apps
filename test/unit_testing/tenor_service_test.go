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


func TestTenorServiceCreate(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)

	request := &schema.TenorReq{
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}

	// Mock AvailableInstalment untuk mengembalikan (*entity.Tenor)(nil), yaitu tipe *entity.Tenor dengan nilai nil
	mockRepo.On("AvailableInstalment", "1", 1).Return((*entity.Tenor)(nil), nil)

	// Mock Save untuk mengembalikan nil yang menunjukkan bahwa penyimpanan berhasil
	mockRepo.On("Save", mock.AnythingOfType("*entity.Tenor")).Return(nil)

	// Membuat service
	mockService := services.NewTenorService(mockRepo)

	// Memanggil fungsi Create
	err := mockService.Create(request)

	// Assertion untuk memastikan tidak ada error
	assert.NoError(t, err)

	// Memastikan semua ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}

func TestTenorServiceDelete(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)
	ID := "1"
	mockService := services.NewTenorService(mockRepo)
	mockRepo.On("Delete", ID).Return(nil)
	err := mockService.Delete(ID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTenorServiceUpdate(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)
	ID := "1"
	request := &schema.TenorReq{

		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}
	mockService := services.NewTenorService(mockRepo)
	mockRepo.On("FindById", ID).Return(&entity.Tenor{}, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)
	err := mockService.Update(ID, request)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func  TestTenorServiceList(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)

	data := &entity.Tenor{
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}

	mockRepo.On("List").Return([]*entity.Tenor{data}, nil)
	mockService := services.NewTenorService(mockRepo)
	response, err := mockService.List()

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, "1", response[0].CustomerID)
	assert.Equal(t, 1, response[0].InstalmentMonth)
	assert.Equal(t, 1000000, response[0].Amount)
}


func TestTenorServiceFindByID(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)
	ID := "1"
	expectedTenor := &entity.Tenor{
		ID: "1",
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}
	mockService := services.NewTenorService(mockRepo)
	mockRepo.On("FindById", ID).Return(expectedTenor, nil)
	response, err := mockService.FindByID(ID)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, "1", response.ID)
	assert.Equal(t, "1", response.CustomerID)
	assert.Equal(t, 1, response.InstalmentMonth)
	assert.Equal(t, 1000000, response.Amount)
}


func TestTenorServiceFindByCustomerID(t *testing.T) {
	mockRepo := new(mocktesting.MockTenorRepository)
	ID := "1"
	expectedTenor := &entity.Tenor{
		ID: "1",
		CustomerID: "1",
		InstalmentMonth: 1,
		Amount: 1000000,
	}
	mockService := services.NewTenorService(mockRepo)
	mockRepo.On("FindByCustomerID", ID).Return([]*entity.Tenor{expectedTenor}, nil)
	response, err := mockService.FindByCustomerID(ID)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, "1", response[0].ID)
	assert.Equal(t, "1", response[0].CustomerID)
	assert.Equal(t, 1, response[0].InstalmentMonth)
	assert.Equal(t, 1000000, response[0].Amount)
}
