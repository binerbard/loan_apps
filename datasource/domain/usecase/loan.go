package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
)



type LoanUsecase interface {
	service.LoanService
}

type loanUsecase struct {
	loanRepository repository.LoanRepository
	loanService service.LoanService
}

func NewULoanUsecase(r repository.LoanRepository, s service.LoanService) LoanUsecase {
	return &loanUsecase{
		loanRepository: r,
		loanService: s,
	}
}

func(u *loanUsecase) Create(credential *schema.LoanReq) error{
	return u.loanService.Create(credential)
}

func(u *loanUsecase) Delete(id string) error{
	return u.loanService.Delete(id)
}

func(u *loanUsecase) Update(id string, credential *schema.LoanReq) error {
	return u.loanService.Update(id, credential)
}

func(u *loanUsecase) FindByID(id string) (*entity.Loan, error){
	return u.loanService.FindByID(id)
}

func (u *loanUsecase) List() ([]*entity.Loan, error){
	return u.loanService.List()
}