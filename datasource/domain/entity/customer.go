package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Customer struct {
	ID         string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"customers_id" validate:"required,uuid"`
	NIK        string         `gorm:"type:varchar(16);unique;not null" json:"nik" validate:"required,numeric,len=16"`
	FullName   string         `gorm:"type:varchar(255);not null" json:"fullname" validate:"required"`
	LegalName  string         `gorm:"type:varchar(255);not null" json:"legalname" validate:"required"`
	BirthDay   time.Time      `gorm:"type:date;not null" json:"birthday" validate:"required"`
	BirthPlace string         `gorm:"type:varchar(255);not null" json:"birthplace" validate:"required"`
	Salary     int            `gorm:"type:int;not null" json:"salary" validate:"required,gte=0"`
	ImgID      string         `gorm:"type:varchar(255);not null" json:"img_ktp" validate:"required,url"`
	ImgSelfie  string         `gorm:"type:varchar(255);not null" json:"selfie" validate:"required,url"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func ValidatorCustomer(customer *Customer) error {
	validate := validator.New()
	err := validate.Struct(customer)
	return err
}