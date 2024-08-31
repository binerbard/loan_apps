package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Tenor struct {
	ID         		string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"customers_id" validate:"required,uuid"`
	CustomerID      string			`gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	Amount     		int            `gorm:"type:int;not null" json:"amount" validate:"required,gte=0"`
	InstalmentMonth	int            `gorm:"type:int;not null" json:"instalment_month" validate:"required,gte=0"`
	CreatedAt  		time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  		time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  		gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func ValidatorTenor(tenor *Tenor) error {
	validate := validator.New()
	err := validate.Struct(tenor)
	return err
}