package entities

import (
	"context"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
)

const (
	CollectionResetPassword = "resetPassword"
)

type ResetPasswordUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	ResetPassword(c context.Context, userID string, resetPassword *forms.ResetPasswordForm) error
	CreateVerificationToken(user *User, secret string, expiry int) (accessToken string, err error)
	SendVerificationEmail(recipientEmail string, encodedToken string, env *bootstrap.Env) (err error)
	GetUserById(c context.Context, userId string) (*User, error)
}

type ResetPasswordRepository interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	ResetPassword(c context.Context, userID string, resetPassword *forms.ResetPasswordForm) error
}
