package controller

import (
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapRegisterFormToUser(form *forms.RegisterUserForm) *entities.User {
	return &entities.User{
		ID:        primitive.NewObjectID(),
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
		Password:  form.Password,
		Phone:     form.Phone,
		Address:   form.Address,
	}
}
