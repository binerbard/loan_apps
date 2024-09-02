package schema

import (
	"github.com/go-playground/validator/v10"
)



type TenorReq struct {
	CustomerID      string			`gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	InstalmentMonth	int            `gorm:"type:int;not null" json:"instalment_month" validate:"required,gte=0"`
	Amount     		int            `gorm:"type:int;not null" json:"amount" validate:"required,gte=0"`
}

func ValidatorTenorReq(tenor *TenorReq) error {
	validate := validator.New()
	err := validate.Struct(tenor)
	return err
}