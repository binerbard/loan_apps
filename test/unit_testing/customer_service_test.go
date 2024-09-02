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

// Test untuk Create
func TestCustomerServiceCreate(t *testing.T) {
    mockRepo := new(mocktesting.MockCustomerRepo)
    request := &schema.CustomerReq{
        NIK: "123", FullName: "test", LegalName: "test", BirthDay: "test", BirthPlace: "test", Salary: 0, ImgID: "test", ImgSelfie: "test",
    }

    mockService := services.NewCustomerService(mockRepo)
    mockRepo.On("Save", mock.Anything).Return(nil)

    err := mockService.Create(request)
    
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

// Test untuk Update
func TestCustomerServiceUpdate(t *testing.T) {
    mockRepo := new(mocktesting.MockCustomerRepo)
    request := &schema.CustomerReq{
        NIK: "123", FullName: "test", LegalName: "test", BirthDay: "test", BirthPlace: "test", Salary: 0, ImgID: "test", ImgSelfie: "test",
    }
    ID := "1"

    mockService := services.NewCustomerService(mockRepo)
    mockRepo.On("FindById", ID).Return(&entity.Customer{}, nil) // Pastikan FindById dipanggil dengan ID dan kembalikan data dummy
    mockRepo.On("Update", mock.Anything).Return(nil)

    err := mockService.Update(ID, request)

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestCustomerServiceDelete(t *testing.T) {
    mockRepo := new(mocktesting.MockCustomerRepo)
    ID := "1"

    // Setup ekspektasi untuk metode Delete
    mockRepo.On("Delete", ID).Return(nil)

    mockService := services.NewCustomerService(mockRepo)
    err := mockService.Delete(ID)

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}


func TestCustomerServiceFindById(t *testing.T) {
    mockRepo := new(mocktesting.MockCustomerRepo)
    ID := "1"
    expectedCustomer := &entity.Customer{
        ID: "1", NIK: "123", FullName: "test", LegalName: "test", BirthDay: "test", BirthPlace: "test", Salary: 0, ImgID: "test", ImgSelfie: "test",
    }

    // Setup ekspektasi untuk metode FindById
    mockRepo.On("FindById", ID).Return(expectedCustomer, nil)

    mockService := services.NewCustomerService(mockRepo)
    customer, err := mockService.FindByID(ID)

    assert.NoError(t, err)
    assert.NotNil(t, customer)
    assert.Equal(t, expectedCustomer.ID, customer.ID)
    assert.Equal(t, expectedCustomer.FullName, customer.FullName)
    assert.Equal(t, expectedCustomer.LegalName, customer.LegalName)
    assert.Equal(t, expectedCustomer.NIK, customer.NIK)
    assert.Equal(t, expectedCustomer.BirthDay, customer.BirthDay)
    assert.Equal(t, expectedCustomer.BirthPlace, customer.BirthPlace)
    assert.Equal(t, expectedCustomer.Salary, customer.Salary)
    assert.Equal(t, expectedCustomer.ImgID, customer.ImgID)
    assert.Equal(t, expectedCustomer.ImgSelfie, customer.ImgSelfie)
    mockRepo.AssertExpectations(t)
}


// Test untuk List
func TestCustomerServiceList(t *testing.T) {
    mockRepo := new(mocktesting.MockCustomerRepo)

    data := &entity.Customer{
        ID: "1", NIK: "123", FullName: "test", LegalName: "test", BirthDay: "test", BirthPlace: "test", Salary: 0, ImgID: "test", ImgSelfie: "test",
    }

    mockRepo.On("List").Return([]*entity.Customer{data}, nil)

    mockService := services.NewCustomerService(mockRepo)
    customers, err := mockService.List()

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)

    // Verifikasi hasil
    assert.Equal(t, "1", customers[0].ID)
    assert.Equal(t, "test", customers[0].FullName)
    assert.Equal(t, "test", customers[0].LegalName)
    assert.Equal(t, "123", customers[0].NIK)
    assert.Equal(t, "test", customers[0].BirthDay)
    assert.Equal(t, "test", customers[0].BirthPlace)
    assert.Equal(t, 0, customers[0].Salary)
    assert.Equal(t, "test", customers[0].ImgID)
    assert.Equal(t, "test", customers[0].ImgSelfie)
}