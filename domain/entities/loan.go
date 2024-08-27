package entities

import (
	"context"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLoan = "loans"
)

type Loan struct {
	LoanID       primitive.ObjectID `bson:"loan_id,omitempty" json:"loanId,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"userId,omitempty"`
	Amount       float64            `bson:"amount" json:"amount" validate:"required,gt=0"`
	InterestRate float64            `bson:"interest_rate" json:"interestRate" validate:"required,gt=0"`
	Term         int                `bson:"term" json:"term" validate:"required,gt=0"` // Term in months
	StartDate    primitive.DateTime `bson:"start_date" json:"startDate" validate:"required"`
	DueDate      primitive.DateTime `bson:"due_date" json:"dueDate" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=active closed delinquent"`
}

type LoanRepository interface {
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, loanID string) (*Loan, error)
	DeleteLoan(ctx context.Context, loanID string) error
	GetLoans(ctx context.Context, limit int64, page int64) (*[]Loan, mongopagination.PaginationData, error)
}

type LoanUseCase interface {
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, loanID string) (*Loan, error)
	GetLoans(ctx context.Context, limit int64, page int64) (*[]Loan, mongopagination.PaginationData, error)
	DeleteLoan(ctx context.Context, loanID string) error
}
