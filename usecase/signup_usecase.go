package usecase

import (
	"context"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	emailutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/email"
)

type signupUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) GetUserById(c context.Context, userId string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserById(ctx, userId)
	return user, err
}
func (su *signupUsecase) ActivateUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.ActivateUser(ctx, userID)
}

func (uu *signupUsecase) FirstUser(c context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	users, err := uu.userRepository.GetAllUsers(ctx)

	if err != nil {
		return false, err
	}

	if len(users) == 0 {
		return true, nil
	}

	return false, nil
}

func (su *signupUsecase) Create(c context.Context, user *entities.User) (*entities.User, error) {

	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.CreateUser(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}

func (su *signupUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {

	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {

	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (su *signupUsecase) CreateVerificationToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {

	return tokenutil.CreateVerificationToken(user, secret, expiry)
}

func (su *signupUsecase) SendVerificationEmail(recipientEmail string, encodedToken string, env *bootstrap.Env) (err error) {

	return emailutil.SendVerificationEmail(recipientEmail, encodedToken, env)
}
