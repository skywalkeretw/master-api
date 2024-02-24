package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/skywalkeretw/master-api/app/v1/function"
	"github.com/skywalkeretw/master-api/app/v1/kubernetes"
	openapi "github.com/skywalkeretw/master-api/app/v1/openAPI"
	"github.com/skywalkeretw/master-api/app/v1/rabbitmq"
)

func ConfigureRoutes(v1 *gin.RouterGroup) {
	v1.Group("/G1").GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})
	v1.GET("/pods", kubernetes.PodHandler)
	v1.GET("/app", kubernetes.GetAllAppsHandler)
	v1.GET("/app/:name", kubernetes.GetAppHandler)
	v1.DELETE("/app/:name", kubernetes.DeleteAppHandler)
	v1.PATCH("/app", kubernetes.UpdateAppHandler)

	v1.GET("/queue", rabbitmq.ListQueuesHandler)
	// v1.GET("/queue/:name", rabbitmq.GetQueueHandler)
	v1.POST("/queue", rabbitmq.CreateQueueHandler)
	v1.GET("/help", openapi.SwaggerCodegenHelpHandler)
	v1.POST("/generateserver", openapi.GenerateServerStubHandler)
	v1.POST("/generateclient", openapi.GenerateClientHandler)

	v1.POST("/createfunction", function.CreateFunctionHandler)
	v1.GET("/function", function.GetFunctionsHandler)
	v1.POST("/generateadaptercode", function.GenerateAdapterCodeHandler)
	// v1.DELETE("/queue", rabbitmq.DeleteQueueHandler)
	// v1.POST("/generateadapter")
	// v1.POST("/deployapp")

}
