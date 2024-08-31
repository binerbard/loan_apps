package schema

import "github.com/go-playground/validator/v10"

type SignInReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// ValidateSignInReq validates the sign in request
func ValidateSignInReq(signIn *SignInReq) error {
	validate := validator.New()
	return validate.Struct(signIn)
}
