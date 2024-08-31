package schema

import (
	"github.com/go-playground/validator/v10"
)

type LoanReq struct {
	CustomerID      string         `gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	TenorID     		string        `gorm:"type:char(36);not null" json:"tenor_id" validate:"required,uuid"`
}

func ValidatorLoanReq(loan *LoanReq) error {
	validate := validator.New()
	err := validate.Struct(loan)
	return err
}