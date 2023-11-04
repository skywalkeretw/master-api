package rabbitmq

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary List RabbitMQ Queues
// @Description List all RabbitMQ Queue available in the system
// @Tags Rabbitmq
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Router /api/v1/rabbitmq [get]
func ListQueuesHandler(ctx *gin.Context) {
	ListQueues()
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type QueueDeclareData struct {
	Name string `json:"name"`
}

// @Summary Create a new queue
// @Description Create a new queue in RabbitMQ.
// @Accept json
// @Produce json
// @Param QueueData body QueueDeclareData true "Queue data"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /createQueue [post]
func CreateQueueHandler(ctx *gin.Context) {
	var queueData QueueDeclareData
	if err := ctx.ShouldBindJSON(&queueData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := CreateQueue(queueData.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": fmt.Sprintf("Queue '%s' created successfully.\n", queueData.Name)})
}
