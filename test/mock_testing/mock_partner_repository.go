package mocktesting

import (
	"loan_apps/datasource/domain/entity"

	"github.com/stretchr/testify/mock"
)


type MockPartnerRepository struct {
	mock.Mock
}

func (p *MockPartnerRepository) Save(partner *entity.Partner) ( error) {
	return p.Called(partner).Error(0)
}

func (p *MockPartnerRepository) List() ([]*entity.Partner, error) {
	args := p.Called()
	return args.Get(0).([]*entity.Partner), args.Error(1)
}

func (p *MockPartnerRepository) FindById(id string) (*entity.Partner, error) {
	args := p.Called(id)
	return args.Get(0).(*entity.Partner), args.Error(1)
}

func (p *MockPartnerRepository) Update(partner *entity.Partner) (error) {
	return p.Called(partner).Error(0)
}

func (p *MockPartnerRepository) Delete(id string) error {
	return p.Called(id).Error(0)
}
