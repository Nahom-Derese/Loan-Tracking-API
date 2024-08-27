package usecase

import (
	"context"
	"log"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
	Env            *bootstrap.Env
}

func NewUserUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.CreateUser(ctx, user)
}

func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByEmail(ctx, email)
}

func (uu *userUsecase) GetUserById(c context.Context, userId string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserById(ctx, userId)
}
func (uu *userUsecase) GetUsers(c context.Context, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	filter := UserFilterOption(userFilter)
	users, meta, err := uu.userRepository.GetUsers(ctx, filter, userFilter)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	// map users to User
	res := make([]entities.User, 0)

	for _, user := range *users {
		res = append(res, user)
	}

	return &res, meta, nil
}
func (uu *userUsecase) UpdateUser(c context.Context, userID string, updatedUser *forms.UpdateUserForm) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.UpdateUser(ctx, userID, updatedUser)
}

func (uu *userUsecase) UpdateUserLoan(c context.Context, userID string, amount float64) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.UpdateLoanAmount(ctx, userID, amount)
}

func (uu *userUsecase) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.DeleteUser(ctx, userID)
}

func (uu *userUsecase) IsUserActive(c context.Context, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.IsUserActive(ctx, userID)
}

func (uu *userUsecase) UpdateUserPassword(c context.Context, userID string, updatePassword *forms.UpdatePasswordForm) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.updateUserPassword(ctx, userID, updatePassword.NewPassword)
}

func (uu *userUsecase) ResetUserPassword(c context.Context, userID string, updatePassword *forms.ResetPasswordForm) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.updateUserPassword(ctx, userID, updatePassword.NewPassword)
}

func UserFilterOption(filter entities.UserFilter) bson.M {

	query := bson.M{
		"$match": bson.M{},
	}
	semiquery := query["$match"].(bson.M)

	// Email filter
	if filter.Email != "" {
		semiquery["email"] = bson.M{"$regex": filter.Email, "$options": "i"}
	}

	// filter.Role
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Active filter
	if filter.Active != "" {
		semiquery["active"] = filter.Active == "true"
	}

	// First name filter
	if filter.FirstName != "" {
		semiquery["first_name"] = bson.M{"$regex": filter.FirstName, "$options": "i"} // case-insensitive search
	}

	// Last name filter
	if filter.LastName != "" {
		semiquery["last_name"] = bson.M{"$regex": filter.LastName, "$options": "i"} // case-insensitive search
	}

	// Is admin filter
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Date range filter
	if !filter.DateFrom.IsZero() && !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{
			"$gte": filter.DateFrom,
			"$lte": filter.DateTo,
		}
	} else if !filter.DateFrom.IsZero() {
		semiquery["created_at"] = bson.M{"$gte": filter.DateFrom}
	} else if !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{"$lte": filter.DateTo}
	}

	log.Println(query)
	return query

}

func (uu *userUsecase) updateUserPassword(ctx context.Context, userID string, newPassword string) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	return uu.userRepository.UpdateUserPassword(ctx, userID, string(encryptedPassword))
}
