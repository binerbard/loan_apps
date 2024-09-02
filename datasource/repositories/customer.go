package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)


type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{DB: db}
}

func (r *CustomerRepositoryImpl) Save(customer *entity.Customer) ( error){
	if err := r.DB.Create(customer).Error ; err != nil {
		return err
	}
	return nil
}
func (r *CustomerRepositoryImpl) Update(customer *entity.Customer) (error){
	if err := r.DB.Save(customer).Error; err != nil {
		return err
	}
	return nil
}
func (r *CustomerRepositoryImpl) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Customer{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryImpl) FindById(id string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	if err := r.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepositoryImpl) List() ([]*entity.Customer, error){
	var Customers []*entity.Customer
	if err := r.DB.Find(&Customers).Error; err != nil {
		return nil, err
	}
	return Customers, nil
}