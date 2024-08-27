package entities

import (
	"context"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"first_name" json:"firstName" validate:"required,min=2,max=100"`
	LastName        string             `bson:"last_name" json:"lastName" validate:"required,min=2,max=100"`
	Email           string             `bson:"email" json:"email" validate:"required,email"`
	Password        string             `bson:"password" json:"-" validate:"required,min=6"`
	Phone           string             `bson:"phone" json:"phone" validate:"required,e164,min=10,max=13"`
	Address         string             `bson:"address" json:"address" validate:"required,min=5,max=200"`
	Active          bool               `bson:"active" json:"active"`
	Role            string             `bson:"role" json:"role" validate:"required,oneof=user admin"`
	TotalLoanAmount float64            `bson:"total_loan_amount" json:"totalLoanAmount" validate:"gte=0"`
	OutstandingDebt float64            `bson:"outstanding_debt" json:"outstandingDebt" validate:"gte=0"`
	Loans           []Loan             `bson:"loans" json:"loans"`
	CreatedAt       primitive.DateTime `bson:"created_at" json:"createdAt"`
	UpdatedAt       primitive.DateTime `bson:"updated_at" json:"updatedAt"`
}

func (u *User) Validate() error {
	validate := bootstrap.GetValidator()
	return validate.Struct(u)
}

type UserFilter struct {
	Email     string
	DateFrom  time.Time
	DateTo    time.Time
	Limit     int64
	Pages     int64
	FirstName string
	LastName  string
	Role      string
	Active    string
	Sort      string
}

type UserUsecase interface {
	CreateUser(c context.Context, user *User) (*User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserById(c context.Context, userId string) (*User, error)
	GetUsers(c context.Context, filter UserFilter) (*[]User, mongopagination.PaginationData, error)
	UpdateUser(c context.Context, userID string, updatedUser *forms.UpdateUserForm) (*User, error)
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)
	UpdateUserPassword(c context.Context, userID string, updatePassword *forms.UpdatePasswordForm) error
	ResetUserPassword(c context.Context, userID string, updatePassword *forms.ResetPasswordForm) error

	// IsOwner(c context.Context) (bool, error)
	// UpdateProfilePicture(c context.Context, userID string, filename string) error
	// PromoteUserToAdmin(c context.Context, userID string) error
	// DemoteAdminToUser(c context.Context, userID string) error
}

type UserRepository interface {
	GetUsers(c context.Context, filter bson.M, userFilter UserFilter) (*[]User, mongopagination.PaginationData, error)
	GetAllUsers(c context.Context) ([]User, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserById(c context.Context, userId string) (*User, error)
	CreateUser(c context.Context, user *User) (*User, error)
	UpdateUser(c context.Context, userID string, updatedUser *forms.UpdateUserForm) (*User, error)
	UpdateRefreshToken(c context.Context, userID string, refreshToken string) error
	UpdateLastLogin(c context.Context, userID string) error
	ActivateUser(c context.Context, userID string) error
	DeleteUser(c context.Context, userID string) error
	IsUserActive(c context.Context, userID string) (bool, error)
	RevokeRefreshToken(c context.Context, userID, refreshToken string) error
	UpdateUserPassword(c context.Context, userID string, newPassword string) error

	// UpdateProfilePicture(c context.Context, userID string, filename string) error
	// PromoteUserToAdmin(c context.Context, userID string) error
	// DemoteAdminToUser(c context.Context, userID string) error
	// GetInactiveUsersForReactivation(c context.Context, emailTreshold primitive.DateTime, deleteTreshold primitive.DateTime) (*[]User, error)
	// DeleteInActiveUser(c context.Context, deleteTreshold primitive.DateTime) error

	RefreshTokenExist(c context.Context, userID, refreshToken string) (bool, error)
}
