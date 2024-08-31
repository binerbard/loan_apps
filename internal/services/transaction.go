package services

import (
	"loan_apps/datasource/domain/entity"
	"loan_apps/datasource/domain/repository"
	"loan_apps/datasource/domain/schema"
	"loan_apps/datasource/domain/service"
	"time"

	"github.com/google/uuid"
)


type TransactionServiceImpl struct {
	transactionData repository.TransactionRepository
}

func NewtTransactionService(transactionRepository repository.TransactionRepository) service.TransactionService {
	return &TransactionServiceImpl{
		transactionData: transactionRepository,
	}
}

func (s *TransactionServiceImpl) Create(credential *schema.TransactionReq) error {
	transaction := entity.Transaction{
		ID: uuid.New().String(),
		CustomerID: credential.CustomerID,
		ContractNumber: credential.ContractNumber,
		OTR: credential.OTR,
		AdminFee: credential.AdminFee,
		InstallmentAmount: credential.InstallmentAmount,
		InterestAmount: credential.InterestAmount,
		AssetName: credential.AssetName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.transactionData.Save(&transaction)

}
func (s *TransactionServiceImpl) Delete(id string) error {
	return s.transactionData.Delete(id)
}

func (s *TransactionServiceImpl) Update(id string, credential *schema.TransactionReq) error {
	transaction, err := s.transactionData.FindById(id)
	if err != nil {
		return err
	}

	transaction.ContractNumber = credential.ContractNumber
	transaction.OTR = credential.OTR
	transaction.AdminFee = credential.AdminFee
	transaction.InstallmentAmount = credential.InstallmentAmount
	transaction.InterestAmount = credential.InterestAmount
	transaction.UpdatedAt = time.Now()

	return s.transactionData.Update(transaction)
}

func (s *TransactionServiceImpl) FindByID(id string) (*entity.Transaction, error) {
	return s.transactionData.FindById(id)
}

func (s *TransactionServiceImpl) List() ([]*entity.Transaction, error){
	return s.transactionData.List()
}