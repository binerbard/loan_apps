package schema

import (
	"github.com/go-playground/validator/v10"
)




type PartnerReq struct {
	Name        string         `gorm:"type:varchar(16);unique;not null" json:"nik" validate:"required"`
	Contact   string         `gorm:"type:varchar(255);not null" json:"fullname" validate:"required"`
}



func ValidatorPartnerReq(partner *PartnerReq) error {
	validate := validator.New()
	err := validate.Struct(partner)
	return err
}