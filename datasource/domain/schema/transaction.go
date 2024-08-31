package schema

import "github.com/go-playground/validator/v10"



type TransactionReq struct {
	CustomerID      string         `gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	ContractNumber  string         `gorm:"type:varchar(255);unique;not null" json:"contract_number" validate:"required"`
	OTR             int            `gorm:"type:int;not null" json:"otr" validate:"required,gte=0"`
	AdminFee        int            `gorm:"type:int;not null" json:"admin_fee" validate:"required,gte=0"`
	InstallmentAmount int          `gorm:"type:int;not null" json:"installment_amount" validate:"required,gte=0"`
	InterestAmount  int            `gorm:"type:int;not null" json:"interest_amount" validate:"required,gte=0"`
	AssetName       string         `gorm:"type:varchar(255);not null" json:"asset_name" validate:"required"`
}

func ValidatorTransactionReq(transaction *TransactionReq) error {
	validate := validator.New()
	err := validate.Struct(transaction)
	return err
}