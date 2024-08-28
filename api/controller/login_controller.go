package controller

import (
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	error_handler "github.com/Nahom-Derese/Loan-Tracking-API/internal/error"
	"github.com/Nahom-Derese/Loan-Tracking-API/internal/logger"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase entities.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request forms.LoginUserForm

	err := c.ShouldBindJSON(&request)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body cannot be empty"})
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

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		// Log failed login attempt
		if err := logger.LogLoginAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		return
	}

	if !user.Active {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotActive))
		// Log failed login attempt
		if err := logger.LogLoginAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		// Log failed login attempt
		if err := logger.LogLoginAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrCredentialsNotValid))
		return
	}

	// Logging
	if err := logger.LogLoginAttempt(user.Email, true); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	err = lc.LoginUsecase.UpdateRefreshToken(c.Request.Context(), user.ID.Hex(), refreshToken)

	loginResponse := entities.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
