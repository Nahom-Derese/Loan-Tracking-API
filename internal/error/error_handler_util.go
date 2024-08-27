package error_handler

import (
	"errors"
	"net/http"

	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CustomErrorResponse(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]custom_error.ValidationErrorResponse, len(ve))
		for i, fe := range ve {
			out[i] = custom_error.ValidationErrorResponse{Field: fe.Field(), Message: "validators.ValidationMessage(fe)"}
		}
		c.JSON(http.StatusBadRequest, custom_error.ErrValidation(out))
		return
	}

	c.JSON(http.StatusBadRequest, custom_error.ErrMessage(err))
}
