package controller

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	error_handler "github.com/Nahom-Derese/Loan-Tracking-API/internal/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// interface for blog controllers
type loansController interface {
	GetLoan() gin.HandlerFunc
	GetLoans() gin.HandlerFunc
	ApplyLoan() gin.HandlerFunc
}

// LoanController is a struct to hold the usecase and env
type LoanController struct {
	LoanUseCase entities.LoanUseCase
	Env         *bootstrap.Env
}

func (pc *LoanController) GetLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		loan, err := pc.LoanUseCase.GetLoanByID(c, id)
		if err != nil {
			c.Error(err)
			return
		}
		if claimUserID != loan.UserID.Hex() && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"loan": loan})
	}
}

func (pc *LoanController) GetLoans() gin.HandlerFunc {
	return func(c *gin.Context) {

		var page int64 = 1
		var limit int64 = 10

		in_page, err := strconv.ParseInt(c.Query("page"), 10, 64)
		if err == nil {
			page = in_page
		}

		in_limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
		if err == nil {
			limit = in_limit
		}

		loans, pagination, err := pc.LoanUseCase.GetLoans(c.Request.Context(), limit, page)

		if err != nil {
			c.Error(err)
			return
		}

		res := entities.PaginatedResponse{
			Data:     loans,
			MetaData: pagination,
		}

		c.JSON(http.StatusOK, res)
	}
}
func (pc *LoanController) ApplyLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimUserID := c.MustGet("x-user-id").(string)

		var request forms.ApplyLoanForm
		if err := c.ShouldBindJSON(&request); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
				return
			}
			error_handler.CustomErrorResponse(c, err)
			return
		}

		if err := request.Validate(); err != nil {
			errorMessages := error_handler.TranslateError(err)
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		userID, err := primitive.ObjectIDFromHex(claimUserID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err})
			return
		}

		loan := entities.Loan{
			LoanID:       primitive.NewObjectID(),
			UserID:       userID,
			Amount:       request.Amount,
			InterestRate: request.InterestRate,
			Term:         request.Term,
			StartDate:    primitive.NewDateTimeFromTime(time.Now()),
			DueDate:      primitive.NewDateTimeFromTime(time.Now()),
		}

		newLoan, err := pc.LoanUseCase.CreateLoan(c, &loan)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"loan": newLoan})
	}
}

func (pc *LoanController) DeleteLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		loanID := c.Param("id")

		err := pc.LoanUseCase.DeleteLoan(c, loanID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
