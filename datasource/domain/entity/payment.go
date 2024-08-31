package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Payment struct {
	ID        		string        	`gorm:"type:char(36);primaryKey;default:uuid()" json:"users_id" validate:"required,uuid"`
	CustomerID		string        	`gorm:"type:char(36);not null" json:"customer_id" validate:"required,uuid"`
	Amount 			int           	`gorm:"type:int;not null" json:"amount" validate:"required,gte=0"`
	PaymentPlatform string 			`gorm:"type:varchar(255);not null" json:"payment_platform" validate:"required"`
	PaymentDate 	time.Time     	`gorm:"type:date;not null" json:"payment_date" validate:"required"`
	CreatedAt 		time.Time     	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 		time.Time     	`gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt	`gorm:"index" json:"deleted_at"`
}

func ValidatorPayment(payment *Payment) error {
	validate := validator.New()
	err := validate.Struct(payment)
	return err
}