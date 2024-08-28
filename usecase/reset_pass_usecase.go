package usecase

import (
	"context"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	emailutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/email"
	"golang.org/x/crypto/bcrypt"
)

type resetPasswordUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewResetPasswordUsecase(userRepo entities.UserRepository, timeout time.Duration) entities.ResetPasswordUsecase {
	return &resetPasswordUsecase{
		userRepository: userRepo,
		contextTimeout: timeout,
	}
}

func (r *resetPasswordUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	user, err := r.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}
func (r *resetPasswordUsecase) ResetPassword(c context.Context, userID string, resetPassword *forms.ResetPasswordForm) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(resetPassword.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	err = r.userRepository.UpdateUserPassword(ctx, userID, string(encryptedPassword))
	return err
}

func (su *resetPasswordUsecase) CreateVerificationToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {

	return tokenutil.CreateVerificationToken(user, secret, expiry)
}

func (su *resetPasswordUsecase) SendVerificationEmail(recipientEmail string, encodedToken string, env *bootstrap.Env) (err error) {

	return emailutil.SendResetPassEmail(recipientEmail, encodedToken, env)
}

func (su *resetPasswordUsecase) GetUserById(c context.Context, userId string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserById(ctx, userId)
	return user, err
}
