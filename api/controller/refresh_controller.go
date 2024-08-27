package controller

import (
	"net/http"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase entities.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request entities.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrMessage(err))
		return
	}

	if valid, err := tokenutil.IsAuthorized(string(request.RefreshToken), rtc.Env.RefreshTokenSecret); !valid || err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrInvalidToken))
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrUserNotFound))
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.ErrMessage(err))
		return
	}

	refreshTokenResponse := entities.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
