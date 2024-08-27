package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	b64 "encoding/base64"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	error_handler "github.com/Nahom-Derese/Loan-Tracking-API/internal/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase entities.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) VerifyEmail(c *gin.Context) {
	Verificationtoken := c.Param("token")
	decodedToken, _ := b64.URLEncoding.DecodeString(Verificationtoken)

	valid, err := tokenutil.IsAuthorized(string(decodedToken), sc.Env.VerificationTokenSecret)

	fmt.Println(string(decodedToken))

	if !valid || err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrInvalidToken))
		return
	}

	claims, err := tokenutil.ExtractUserClaimsFromToken(string(decodedToken), sc.Env.VerificationTokenSecret)
	userID := claims["id"].(string)

	fmt.Println("userID")
	fmt.Println(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}
	user, err := sc.SignupUsecase.GetUserById(context.TODO(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}
	if user.Active {
		c.JSON(http.StatusConflict, custom_error.ErrMessage(custom_error.ErrAlreadyVerified))
		return
	}

	err = sc.SignupUsecase.ActivateUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})

}
func (sc *SignupController) Register(c *gin.Context) {
	var request forms.RegisterUserForm

	err := c.ShouldBindJSON(&request)
	if err != nil {
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

	first, err := sc.SignupUsecase.FirstUser(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	var role = "user"

	if first {
		role = "admin"
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, custom_error.ErrMessage(custom_error.ErrUserAlreadyExists))
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.Error(err)
		return
	}

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:        primitive.NewObjectID(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Active:    false,
		Role:      role,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": error_handler.TranslateError(err)})
		errorMessages := error_handler.TranslateError(err)
		log.Println(errorMessages)
		return
	}

	VerificationToken, err := sc.SignupUsecase.CreateVerificationToken(&user, sc.Env.VerificationTokenSecret, sc.Env.VerificationTokenExpiryMin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	_, err = sc.SignupUsecase.Create(c, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	//send email
	encodedToken := b64.URLEncoding.EncodeToString([]byte(VerificationToken))
	err = sc.SignupUsecase.SendVerificationEmail(user.Email, encodedToken, sc.Env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
	}

	c.JSON(http.StatusCreated, gin.H{"message": "email sent successfully, please verify your email"})

}
