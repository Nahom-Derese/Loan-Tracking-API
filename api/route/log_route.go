package route

import (
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/api/controller"
	"github.com/Nahom-Derese/Loan-Tracking-API/bootstrap"
	"github.com/gin-gonic/gin"
)

func NewLogRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/logs", controller.AdminLogsHandler)
}
