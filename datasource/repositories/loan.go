package repositories

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"

	"gorm.io/gorm"
)


type LoanRepositoryImpl struct {
	DB *gorm.DB
}

func NewLoanRepository(db *gorm.DB) repository.LoanRepository {
	return &LoanRepositoryImpl{DB: db}
}



func (r *LoanRepositoryImpl) Save(loan *entity.Loan, balance *entity.Balance) ( error){
	tx := r.DB.Begin()
	if err := tx.Create(loan).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(balance).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (r *LoanRepositoryImpl) Update(loan *entity.Loan) (error){
	if err := r.DB.Save(loan).Error; err != nil {
		return err
	}
	return nil
}
func (r *LoanRepositoryImpl) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Loan{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *LoanRepositoryImpl) FindById(id string) (*entity.Loan, error) {
	Loan := &entity.Loan{}
	if err := r.DB.Where("id = ?", id).First(Loan).Error; err != nil {
		return nil, err
	}
	return Loan, nil
}

func (r *LoanRepositoryImpl) List() ([]*entity.Loan, error){
	var Loans []*entity.Loan
	if err := r.DB.Find(&Loans).Error; err != nil {
		return nil, err
	}
	return Loans, nil
}

func (r *LoanRepositoryImpl) FindByCustomerID(customer_id string) ([]*entity.Loan, error) {
	var Loans []*entity.Loan
	if err := r.DB.Where("customer_id = ?", customer_id).Find(&Loans).Error; err != nil {
		return nil, err
	}
	return Loans, nil
}