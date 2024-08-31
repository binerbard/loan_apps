package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"transaction_id" validate:"required,uuid"`
	CustomerID      string         `gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	ContractNumber  string         `gorm:"type:varchar(255);unique;not null" json:"contract_number" validate:"required"`
	OTR             int            `gorm:"type:int;not null" json:"otr" validate:"required,gte=0"`
	AdminFee        int            `gorm:"type:int;not null" json:"admin_fee" validate:"required,gte=0"`
	InstallmentAmount int          `gorm:"type:int;not null" json:"installment_amount" validate:"required,gte=0"`
	InterestAmount  int            `gorm:"type:int;not null" json:"interest_amount" validate:"required,gte=0"`
	AssetName       string         `gorm:"type:varchar(255);not null" json:"asset_name" validate:"required"`
	CreatedAt       time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"type:timestamp;index" json:"deleted_at"`
}

func ValidatorTransaction(transaction *Transaction) error {
	validate := validator.New()
	err := validate.Struct(transaction)
	return err
}