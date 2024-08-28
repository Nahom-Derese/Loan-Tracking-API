package entities

import (
	"context"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
)

type SignupResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) (*User, error)
	FirstUser(c context.Context) (bool, error)
	ActivateUser(c context.Context, userID string) error
	GetUserById(c context.Context, userId string) (*User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateVerificationToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	SendVerificationEmail(recipientEmail string, encodedToken string, env *bootstrap.Env) (err error)
}
