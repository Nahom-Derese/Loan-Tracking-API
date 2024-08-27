package controller

import (
	b64 "encoding/base64"
	"io"
	"net/http"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	error_handler "github.com/Nahom-Derese/Loan-Tracking-API/internal/error"
	"github.com/Nahom-Derese/Loan-Tracking-API/internal/logger"
	"github.com/gin-gonic/gin"
)

type ResetPasswordController struct {
	ResetPasswordUsecase entities.ResetPasswordUsecase
	Env                  *bootstrap.Env
}

func (rc *ResetPasswordController) ForgotPassword(c *gin.Context) {
	var req forms.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := rc.ResetPasswordUsecase.GetUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	VerificationToken, err := rc.ResetPasswordUsecase.CreateVerificationToken(&user, rc.Env.VerificationTokenSecret, rc.Env.VerificationTokenExpiryMin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	//send email
	encodedToken := b64.URLEncoding.EncodeToString([]byte(VerificationToken))
	err = rc.ResetPasswordUsecase.SendVerificationEmail(user.Email, encodedToken, rc.Env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reset password guide sent successfully, please check your email"})
}

func (rc *ResetPasswordController) ResetPassword(c *gin.Context) {
	var request forms.ResetPasswordForm
	Verificationtoken := c.Param("token")
	decodedToken, _ := b64.URLEncoding.DecodeString(Verificationtoken)

	valid, err := tokenutil.IsAuthorized(string(decodedToken), rc.Env.VerificationTokenSecret)

	if !valid || err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrInvalidToken))
		return
	}

	claims, err := tokenutil.ExtractUserClaimsFromToken(string(decodedToken), rc.Env.VerificationTokenSecret)
	userID := claims["id"].(string)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	user, err := rc.ResetPasswordUsecase.GetUserById(c.Request.Context(), userID)

	if err != nil {
		// Log failed login attempt
		if err := logger.ResetPasswordAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		// Log failed login attempt
		if err := logger.ResetPasswordAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, custom_error.ErrMessage(custom_error.EreInvalidRequestBody))
			return
		}
		error_handler.CustomErrorResponse(c, err)
		return
	}

	if err := request.Validate(); err != nil {
		// Log failed login attempt
		if err := logger.ResetPasswordAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		errorMessages := error_handler.TranslateError(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	err = rc.ResetPasswordUsecase.ResetPassword(c.Request.Context(), user.ID.Hex(), &request)

	if err != nil {
		// Log failed login attempt
		if err := logger.ResetPasswordAttempt(user.Email, false); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
			return
		}
		c.Error(err)
		return
	}

	// Logging
	if err := logger.ResetPasswordAttempt(user.Email, true); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
