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

	publicRouter := gin.Group("")

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewPublicResetPasswordRouter(env, timeout, db, publicRouter)

	// Static files
	// NewPublicFileRouter(env, publicRouter)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	// All Protected APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	// NewLoansRouter(env, timeout, db, protectedRouter)
}
