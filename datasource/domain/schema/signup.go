package schema

import "github.com/go-playground/validator/v10"

// SignUpReq - schema for sign up
type SignUpReq struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

// ValidateSignUp - validate sign up request
func ValidateSignUpReq(signUp *SignUpReq) error {
	validate := validator.New()
	return validate.Struct(signUp)
}
