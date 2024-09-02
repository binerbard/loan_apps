package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)



type Partner struct {
	ID         string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"id" validate:"required,uuid"`
	Name        string         `gorm:"type:varchar(16);unique;not null" json:"nik" validate:"required"`
	Contact   string         `gorm:"type:varchar(255);not null" json:"fullname" validate:"required"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}



func ValidatorPartner(partner *Partner) error {
	validate := validator.New()
	err := validate.Struct(partner)
	return err
}