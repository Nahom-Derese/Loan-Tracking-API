package usecase

import (
	"context"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"
)

type loanUseCase struct {
	loanRepository entities.LoanRepository
	contextTimeout time.Duration
}

func NewLoanUsecase(loanRepository entities.LoanRepository, timeout time.Duration) entities.LoanUseCase {
	return &loanUseCase{
		loanRepository: loanRepository,
		contextTimeout: timeout,
	}
}

// CreateLoan handles the business logic for creating a new loan.
func (uc *loanUseCase) CreateLoan(ctx context.Context, loan *entities.Loan) (*entities.Loan, error) {
	// Perform any necessary business logic or validation here
	// For example, check if the loan amount is within acceptable limits

	if loan.Amount <= 0 {
		return nil, custom_error.ErrInvalidLoanAmount
	}

	// Call the repository to create the loan
	createdLoan, err := uc.loanRepository.CreateLoan(ctx, loan)
	if err != nil {
		return nil, err
	}

	return createdLoan, nil
}

// GetLoanByID handles the business logic for retrieving a loan by its ID.
func (uc *loanUseCase) GetLoanByID(ctx context.Context, loanID string) (*entities.Loan, error) {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	loan, err := uc.loanRepository.GetLoanByID(c, loanID)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (uc *loanUseCase) GetLoans(ctx context.Context, limit int64, page int64) (*[]entities.Loan, mongopagination.PaginationData, error) {

	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	return uc.loanRepository.GetLoans(c, limit, page)

}

// DeleteLoan handles the business logic for deleting a loan by its ID.
func (uc *loanUseCase) DeleteLoan(ctx context.Context, loanID string) error {
	c, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.loanRepository.DeleteLoan(c, loanID)
	if err != nil {
		return err
	}

	return nil
}
