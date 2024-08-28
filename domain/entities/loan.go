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
	StartDate    primitive.DateTime `bson:"start_date" json:"startDate" validate:"required"`
	DueDate      primitive.DateTime `bson:"due_date" json:"dueDate" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=pending approved rejected"`
	Purpose      string             `bson:"purpose" json:"purpose" validate:"required"`
}

type LoanRepository interface {
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, loanID string) (*Loan, error)
	DeleteLoan(ctx context.Context, loanID string) error
	GetLoans(ctx context.Context, limit int64, page int64) (*[]Loan, mongopagination.PaginationData, error)
	RejectLoan(ctx context.Context, id string) error
	AcceptLoan(ctx context.Context, id string) error
}

type LoanUseCase interface {
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, loanID string) (*Loan, error)
	GetLoans(ctx context.Context, limit int64, page int64) (*[]Loan, mongopagination.PaginationData, error)
	DeleteLoan(ctx context.Context, loanID string) error
	AcceptLoan(ctx context.Context, loanID string) error
	RejectLoan(ctx context.Context, loanID string) error
}
