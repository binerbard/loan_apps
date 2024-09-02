package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Balance struct {
	ID        	string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"id" validate:"required,uuid"`
	CustomerID	string         `gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	Balance 	int            `gorm:"type:int;not null" json:"balance" validate:"required,gte=0"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func ValidatorBalance(balance *Balance) error {
	validate := validator.New()
	err := validate.Struct(balance)
	return err
}