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



func (r *LoanRepositoryImpl) Save(loan *entity.Loan) ( error){
	if err := r.DB.Create(loan).Error ; err != nil {
		return err
	}
	return nil
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