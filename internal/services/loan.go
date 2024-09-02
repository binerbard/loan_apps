package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
	"time"

	"github.com/google/uuid"
)


type LoanServiceImpl struct {
	loanData repository.LoanRepository
	tenorData repository.TenorRepository
}

func NewtLoanService(loanRepository repository.LoanRepository, tenorRepository repository.TenorRepository) service.LoanService {
	return &LoanServiceImpl{
		loanData: loanRepository,
		tenorData: tenorRepository,
	}
}

func (s *LoanServiceImpl) Create(credential *schema.LoanReq) error {
	
	acceptTenor,err := s.tenorData.FindById(credential.TenorID)
	if err != nil {
		return err
	}

	balance := entity.Balance{
		ID: uuid.New().String(),
		CustomerID: credential.CustomerID,
		Balance: acceptTenor.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	loan := entity.Loan{
		ID: uuid.New().String(),
		CustomerID: credential.CustomerID,
		Amount: acceptTenor.Amount,
		InstalmentMonth: acceptTenor.InstalmentMonth,
		LoanDate: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.loanData.Save(&loan, &balance)

}
func (s *LoanServiceImpl) Delete(id string) error {
	return s.loanData.Delete(id)
}

func (s *LoanServiceImpl) Update(id string, credential *schema.LoanReq) error {
	loan, err := s.loanData.FindById(id)
	if err != nil {
		return err
	}

	acceptTenor,err := s.tenorData.FindById(credential.TenorID)
	if err != nil {
		return err
	}


	loan.Amount = acceptTenor.Amount
	loan.InstalmentMonth = acceptTenor.InstalmentMonth
	loan.UpdatedAt = time.Now()

	return s.loanData.Update(loan)
}

func (s *LoanServiceImpl) FindByID(id string) (*entity.Loan, error) {
	return s.loanData.FindById(id)
}

func (s *LoanServiceImpl) List() ([]*entity.Loan, error){
	return s.loanData.List()
}

func (s *LoanServiceImpl) FindByCustomerID(customer_id string) ([]*entity.Loan, error){
	return s.loanData.FindByCustomerID(customer_id)
}