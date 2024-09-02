package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Dealer struct {
	ID         		string         `gorm:"type:char(36);primaryKey;default:uuid()" json:"id" validate:"required,uuid"`
	TransactionID      string			`gorm:"type:char(36);not null" json:"transaction_id" validate:"required,uuid"`
	PartnerID      string			`gorm:"type:char(36);not null" json:"partner_id" validate:"required,uuid"`
	CreatedAt  		time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  		time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  		gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func ValidatorDelaer(delaer *Dealer) error {
	validate := validator.New()
	err := validate.Struct(delaer)
	return err
}