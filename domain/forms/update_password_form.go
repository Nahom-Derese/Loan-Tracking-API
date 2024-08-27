package forms

import "github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"

type UpdatePasswordForm struct {
	OldPassword string `json:"oldPassword" validate:"required,min=6"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type ResetPasswordForm struct {
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (u *UpdatePasswordForm) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}

func (u *ResetPasswordForm) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}
