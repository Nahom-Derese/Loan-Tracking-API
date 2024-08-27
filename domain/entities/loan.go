package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct {
	LoanID       primitive.ObjectID `bson:"loan_id,omitempty" json:"loanId,omitempty"`
	Amount       float64            `bson:"amount" json:"amount" validate:"required,gt=0"`
	InterestRate float64            `bson:"interest_rate" json:"interestRate" validate:"required,gt=0"`
	Term         int                `bson:"term" json:"term" validate:"required,gt=0"` // Term in months
	StartDate    int64              `bson:"start_date" json:"startDate" validate:"required"`
	DueDate      int64              `bson:"due_date" json:"dueDate" validate:"required"`
	Status       string             `bson:"status" json:"status" validate:"required,oneof=active closed delinquent"`
}
