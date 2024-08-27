package forms

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ApplyLoanForm represents the form used to apply for a loan.
type ApplyLoanForm struct {
	Amount       float64            `json:"amount" bson:"amount" validate:"required,gt=0"`
	InterestRate float64            `json:"interestRate" bson:"interest_rate" validate:"required,gt=0"`
	Term         int                `json:"term" bson:"term" validate:"required,gt=0"` // Term in months
	StartDate    primitive.DateTime `json:"startDate" bson:"start_date" validate:"required"`
	DueDate      primitive.DateTime `json:"dueDate" bson:"due_date" validate:"required"`
	UserID       string             `json:"user_id" bson:"user_id" validate:"required"` // The ID of the user applying for the loan
	Purpose      string             `json:"purpose" bson:"purpose" validate:"required"` // Reason for the loan
}

// ValidateApplyLoanForm validates the ApplyLoanForm fields.
func (f *ApplyLoanForm) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
