package forms

type UpdatePasswordForm struct {
	OldPassword string `json:"oldPassword" validate:"required,min=6"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type ResetPasswordForm struct {
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}
