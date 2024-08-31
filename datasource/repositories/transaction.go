package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)


type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{DB: db}
}



func (r *TransactionRepositoryImpl) Save(transaction *entity.Transaction) ( error){
	if err := r.DB.Create(transaction).Error ; err != nil {
		return err
	}
	return nil
}
func (r *TransactionRepositoryImpl) Update(transaction *entity.Transaction) (error){
	if err := r.DB.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}
func (r *TransactionRepositoryImpl) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Transaction{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) FindById(id string) (*entity.Transaction, error) {
	Transaction := &entity.Transaction{}
	if err := r.DB.Where("id = ?", id).First(Transaction).Error; err != nil {
		return nil, err
	}
	return Transaction, nil
}

func (r *TransactionRepositoryImpl) List() ([]*entity.Transaction, error){
	var Transactions []*entity.Transaction
	if err := r.DB.Find(&Transactions).Error; err != nil {
		return nil, err
	}
	return Transactions, nil
}