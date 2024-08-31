package usecase

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
)



type TransactionUsecase interface {
	service.TransactionService
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	transactionService service.TransactionService
}

func NewUTransactionUsecase(r repository.TransactionRepository, s service.TransactionService) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: r,
		transactionService: s,
	}
}

func(u *transactionUsecase) Create(credential *schema.TransactionReq) error{
	return u.transactionService.Create(credential)
}

func(u *transactionUsecase) Delete(id string) error{
	return u.transactionService.Delete(id)
}

func(u *transactionUsecase) Update(id string, credential *schema.TransactionReq) error {
	return u.transactionService.Update(id, credential)
}

func(u *transactionUsecase) FindByID(id string) (*entity.Transaction, error){
	return u.transactionService.FindByID(id)
}

func (u *transactionUsecase) List() ([]*entity.Transaction, error){
	return u.transactionService.List()
}