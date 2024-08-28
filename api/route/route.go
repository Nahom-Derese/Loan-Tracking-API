package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/api/middleware"
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {

	// Error handling
	gin.Use(middleware.ErrorHandlerMiddleware())

	usersRouter := gin.Group("users")

	// All Public APIs
	NewSignupRouter(env, timeout, db, usersRouter)
	NewLoginRouter(env, timeout, db, usersRouter)
	NewRefreshTokenRouter(env, timeout, db, usersRouter)
	NewResetPasswordRouter(env, timeout, db, usersRouter)

	protectedUserRouter := gin.Group("users")
	protectedUserRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewProfileRouter(env, timeout, db, protectedUserRouter)

	protectedLoanRouter := gin.Group("loans")
	protectedLoanRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewLoanRouter(env, timeout, db, protectedLoanRouter)

	adminRouter := gin.Group("admin")
	adminRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	adminRouter.Use(middleware.AdminMiddleware())

	NewProfileRouter(env, timeout, db, adminRouter)
	NewLogRouter(env, timeout, adminRouter)
	NewAdminLoanRouter(env, timeout, db, adminRouter)
}
