package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewProfileRouter is a function that defines all the routes for the profile
func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(*db, entities.CollectionUser)
	// pc := controller.ProfileController{
	// 	UserUsecase: usecase.NewUserUsecase(ur, timeout),
	// 	Env:         env,
	// }

	// group.GET("/profiles", middleware.AdminMiddleware(), pc.GetProfiles())
	// group.GET("/profiles/:id", pc.GetProfile())
	// group.PUT("/profiles/:id", pc.UpdateProfile())
	// group.DELETE("/profiles/:id", pc.DeleteProfile())
	// group.POST("/profiles/", pc.ChangePassword())

}
