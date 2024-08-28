package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/api/controller"
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	"github.com/Nahom-Derese/Loan-Tracking-API/repository"
	"github.com/Nahom-Derese/Loan-Tracking-API/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewResetPasswordRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, entities.CollectionUser)

	sc := controller.ResetPasswordController{
		ResetPasswordUsecase: usecase.NewResetPasswordUsecase(ur, timeout),
		Env:                  env,
	}
	group.POST("/password-update/verify/:token", sc.ResetPassword)
	group.POST("/password-reset", sc.ForgotPassword)
}
