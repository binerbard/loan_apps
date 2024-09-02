package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Loan struct {
	ID         		string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"id" validate:"required,uuid"`
	CustomerID      string         `gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	Amount     		int            `gorm:"type:int;not null" json:"amount" validate:"required,gte=0"`
	InstalmentMonth	int            `gorm:"type:int;not null" json:"instalment_month" validate:"required,gte=0"`
	LoanDate		time.Time		`gorm:"type:date;not null" json:"birthday" validate:"required"`
	CreatedAt  		time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  		time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  		gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func ValidatorLoan(loan *Loan) error {
	validate := validator.New()
	err := validate.Struct(loan)
	return err
}