package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)

type BalanceRepo struct {
	DB *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) repository.BalanceRepository {
	return &BalanceRepo{DB: db}
}

func (r *BalanceRepo) Save(balance *entity.Balance) ( error){
	if err := r.DB.Create(balance).Error; err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) Update(balance *entity.Balance) (error){
	if err := r.DB.Save(balance).Error; err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Balance{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) FindById(id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	if err := r.DB.Where("id = ?", id).First(balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}

func (r *BalanceRepo) List() ([]*entity.Balance, error) {
	var balances []*entity.Balance
	if err := r.DB.Find(&balances).Error; err != nil {
		return nil, err
	}
	return balances, nil
}

func (r *BalanceRepo) FindByCustomerID(id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	if err := r.DB.Where("customer_id = ?", id).First(balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}

