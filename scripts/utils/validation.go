package utils

import (
	"github.com/go-playground/validator/v10"
)

// VerifyRequest : Function for verify request
// verify any
func VerifyRequest(verify any) error {

	validate := validator.New(validator.WithRequiredStructEnabled())
	// Validate Struct
	if err := validate.Struct(verify); err != nil {
		return err
	}
	return nil
}
