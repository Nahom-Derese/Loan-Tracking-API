package forms

import (
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
)

type RegisterUserForm struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=100"`
	LastName  string `json:"lastName" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Phone     string `json:"phone" validate:"required,min=10,max=13,e164"`
	Address   string `json:"address" validate:"required,min=5,max=200"`
}

type UpdateUserForm struct {
	FirstName string `bson:"firstName,omitempty" json:"firstName" validate:"min=2,max=100"`
	LastName  string `bson:"lastName,omitempty" json:"lastName" validate:"min=2,max=100"`
	Email     string `bson:"email,omitempty" json:"email" validate:"email"`
	Password  string `bson:"password,omitempty" json:"password" validate:"min=6"`
	Phone     string `bson:"phone,omitempty" json:"phone" validate:"min=10,max=13,e164"`
	Address   string `bson:"address,omitempty" json:"address" validate:"min=5,max=200"`
}

func (u *RegisterUserForm) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}

func (u *UpdateUserForm) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}
