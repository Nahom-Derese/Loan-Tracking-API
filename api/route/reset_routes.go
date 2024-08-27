package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPublicResetPasswordRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewResetPasswordRepository(*db, entities.CollectionUser, entities.CollectionResetPassword)
	// sc := controller.ResetPasswordController{
	// 	ResetPasswordUsecase: usecase.NewResetPasswordUsecase(ur, timeout),
	// 	Env:                  env,
	// }
	// group.POST("/password-update", sc.ResetPassword)
	// group.POST("/password-reset", sc.ForgotPassword)
}
