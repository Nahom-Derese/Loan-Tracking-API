package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/api/middleware"
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, gin *gin.Engine) {

	// Logging
	gin.Use(middleware.RequestLogger())

	// Error handling
	gin.Use(middleware.ErrorHandlerMiddleware())

	usersRouter := gin.Group("users")

	// All Public APIs
	NewSignupRouter(env, timeout, db, usersRouter)
	NewLoginRouter(env, timeout, db, usersRouter)
	NewRefreshTokenRouter(env, timeout, db, usersRouter)
	NewPublicResetPasswordRouter(env, timeout, db, usersRouter)

	protectedUserRouter := gin.Group("users")
	protectedUserRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewProfileRouter(env, timeout, db, protectedUserRouter)

	adminRouter := gin.Group("admin")
	adminRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	adminRouter.Use(middleware.AdminMiddleware())

	// NewLoansRouter(env, timeout, db, protectedRouter)
}
