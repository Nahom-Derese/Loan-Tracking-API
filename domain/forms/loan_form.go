package forms

import (
	"github.com/go-playground/validator/v10"
)

// ApplyLoanForm represents the form used to apply for a loan.
type ApplyLoanForm struct {
	Amount  float64 `json:"amount" bson:"amount" validate:"required,gt=0"`
	DueDate string  `json:"dueDate" bson:"due_date" validate:"required"`
	Purpose string  `json:"purpose" bson:"purpose" validate:"required"` // Reason for the loan
}

// ValidateApplyLoanForm validates the ApplyLoanForm fields.
func (f *ApplyLoanForm) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
