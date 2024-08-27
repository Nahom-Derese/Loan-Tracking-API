package controllers

import (
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapRegisterFormToUser(form *forms.RegisterUserForm) *entites.User {
	return &entites.User{
		ID:        primitive.NewObjectID(),
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
		Password:  form.Password, // Make sure to hash the password
		Phone:     form.Phone,
		Address:   form.Address,
	}
}
