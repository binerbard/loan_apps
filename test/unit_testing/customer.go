package schema

import (
	"github.com/go-playground/validator/v10"
)

type CustomerReq struct {
	NIK        string         `gorm:"type:varchar(16);unique;not null" json:"nik" validate:"required,numeric,len=16"`
	FullName   string         `gorm:"type:varchar(255);not null" json:"fullname" validate:"required"`
	LegalName  string         `gorm:"type:varchar(255);not null" json:"legalname" validate:"required"`
	BirthDay   string      `gorm:"type:date;not null" json:"birthday" validate:"required"`
	BirthPlace string         `gorm:"type:varchar(255);not null" json:"birthplace" validate:"required"`
	Salary     int            `gorm:"type:int;not null" json:"salary" validate:"required,gte=0"`
	ImgID      string         `gorm:"type:varchar(255);not null" json:"img_ktp"`
	ImgSelfie  string         `gorm:"type:varchar(255);not null" json:"selfie"`
}

func ValidatorCustomerReq(customer *CustomerReq) error {
	validate := validator.New()
	err := validate.Struct(customer)
	return err
}