package forms

import "github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"

type LoginUserForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *LoginUserForm) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}
