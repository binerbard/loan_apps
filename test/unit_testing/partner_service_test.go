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

func TestPartnerServiceCreate(t *testing.T) {

	mockPartnerRepository := new(mocktesting.MockPartnerRepository)

	partnerService := services.NewPartnerService(mockPartnerRepository)

	partner := &schema.PartnerReq{
		Name: "test",
		Contact: "test",
	}
	mockPartnerRepository.On("Save", mock.Anything).Return(nil)

	err := partnerService.Create(partner)

	assert.NoError(t, err)

	mockPartnerRepository.AssertExpectations(t)

}

func TestPartnerServiceUpdate(t *testing.T) {
	mockPartnerRepository := new(mocktesting.MockPartnerRepository)

	partnerService := services.NewPartnerService(mockPartnerRepository)

	partner := &schema.PartnerReq{
		Name: "test",
		Contact: "test",
	}
	ID := "1"

	mockPartnerRepository.On("FindById", ID).Return(&entity.Partner{}, nil)
	mockPartnerRepository.On("Update", mock.Anything).Return(nil)

	err := partnerService.Update("1", partner)

	assert.NoError(t, err)

	mockPartnerRepository.AssertExpectations(t)

}

func TestPartnerServiceDelete(t *testing.T) {
	mockPartnerRepository := new(mocktesting.MockPartnerRepository)

	partnerService := services.NewPartnerService(mockPartnerRepository)

	ID := "1"

	mockPartnerRepository.On("Delete", ID).Return(nil)

	err := partnerService.Delete(ID)

	assert.NoError(t, err)

	mockPartnerRepository.AssertExpectations(t)

}

func TestPartnerServiceList(t *testing.T) {
	mockPartnerRepository := new(mocktesting.MockPartnerRepository)

	list := []*entity.Partner{
		{
			ID: "1",
			Name: "test",
			Contact: "test",
		},
	}
	partnerService := services.NewPartnerService(mockPartnerRepository)

	mockPartnerRepository.On("List").Return(list, nil)

	partners, err := partnerService.List()

	assert.NoError(t, err)

	mockPartnerRepository.AssertExpectations(t)

	assert.Equal(t, list, partners)

}

func TestPartnerServiceFindById(t *testing.T) {
	mockPartnerRepository := new(mocktesting.MockPartnerRepository)

	partnerService := services.NewPartnerService(mockPartnerRepository)

	ID := "1"

	expectedPartner := &entity.Partner{
		ID: "1",
		Name: "test",
		Contact: "test",
	}

	mockPartnerRepository.On("FindById", ID).Return(expectedPartner, nil)

	partner, err := partnerService.FindByID(ID)

	assert.NoError(t, err)

	mockPartnerRepository.AssertExpectations(t)

	assert.Equal(t, expectedPartner, partner)

}

