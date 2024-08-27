package controller

import (
	"context"
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
	"golang.org/x/crypto/bcrypt"
)

// interface for blog controllers
type profileController interface {
	GetProfile() gin.HandlerFunc
	UpdateProfile() gin.HandlerFunc
	ChangePassword() gin.HandlerFunc
}

// ProfileController is a struct to hold the usecase and env
type ProfileController struct {
	UserUsecase entities.UserUsecase
	Env         *bootstrap.Env
}

func (pc *ProfileController) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		user, err := pc.UserUsecase.GetUserById(c, claimUserID)
		if err != nil {
			c.Error(err)
			return
		}
		if claimUserID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func (pc *ProfileController) GetProfiles() gin.HandlerFunc {
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

		dateFrom, _ := time.Parse(time.RFC3339, c.Query("date_from"))
		dateTo, _ := time.Parse(time.RFC3339, c.Query("date_to"))

		var userFilter entities.UserFilter

		userFilter = entities.UserFilter{
			Email:     c.Query("email"),
			FirstName: c.Query("first_name"),
			LastName:  c.Query("last_name"),
			Role:      c.Query("role"),
			Active:    c.Query("active"),
			DateFrom:  dateFrom,
			DateTo:    dateTo,
			Limit:     limit,
			Pages:     page,
		}

		users, pagination, err := pc.UserUsecase.GetUsers(context.Background(), userFilter)

		if err != nil {
			c.Error(err)
			return
		}

		res := entities.PaginatedResponse{
			Data:     users,
			MetaData: pagination,
		}

		c.JSON(http.StatusOK, res)
	}
}

func (pc *ProfileController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request forms.UpdatePasswordForm
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

		userID, exists := c.Get("x-user-id")
		if !exists {
			c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUnauthorized))
			return
		}

		// Now you can use userID which is of type interface{}
		userIDStr, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(custom_error.ErrInvalidID))
			return
		}

		user, err := pc.UserUsecase.GetUserById(c, userIDStr)
		if err != nil {
			c.Error(err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
		if err != nil {
			c.Error(err)
			return
		}

		err = pc.UserUsecase.UpdateUserPassword(c, userIDStr, &request)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
	}
}
func (pc *ProfileController) UpdateProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		var request forms.UpdateUserForm
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

		if userID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUnauthorized))
			return
		}

		updatedUser, err := pc.UserUsecase.UpdateUser(c, userID, &request)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": updatedUser})
	}
}

func (pc *ProfileController) DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		claimUserID := c.MustGet("x-user-id").(string)
		role := c.MustGet("x-user-role").(string)

		if userID != claimUserID && role != "admin" {
			c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUnauthorized))
			return
		}

		err := pc.UserUsecase.DeleteUser(c, userID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
