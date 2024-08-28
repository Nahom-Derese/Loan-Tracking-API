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

// NewProfileRouter is a function that defines all the routes for the profile
func NewLoanRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	lr := repository.NewLoanRepository(*db, entities.CollectionLoan)
	ur := repository.NewUserRepository(*db, entities.CollectionUser)
	pc := controller.LoanController{
		LoanUseCase: usecase.NewLoanUsecase(lr, timeout),
		UserUseCase: usecase.NewUserUsecase(ur, timeout),
		Env:         env,
	}

	group.GET(":id", pc.GetLoan())
	group.POST("", pc.ApplyLoan())

}

func NewAdminLoanRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	lr := repository.NewLoanRepository(*db, entities.CollectionLoan)
	pc := controller.LoanController{
		LoanUseCase: usecase.NewLoanUsecase(lr, timeout),
		Env:         env,
	}

	group.GET("/loans", pc.GetLoans())
	group.PATCH("loans/:id/:status", pc.UpdateLoan())
	group.DELETE("loans/:id", pc.DeleteLoan())

}
