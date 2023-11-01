package kubernets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Summary List Kubernetes Pods
// @Description List all Kubernetes pods in the cluster
// @Tags Deployments
// @Accept json
// @Produce json
// @Success 200 {array} PodInfo
// @Router /api/v1/pods [get]
func PodHandler(ctx *gin.Context) {
	p, _ := GetKubernetesPods()
	ctx.JSON(http.StatusOK, p)
}

// @Summary List Kubernetes Deployments
// @Description List all Kubernetes deployments in the cluster
// @Tags Deployments
// @Accept json
// @Produce json
// @Success 200 {array} DeploymentInfo
// @Router /api/v1/app [get]
func GetAllAppsHandler(ctx *gin.Context) {
	d, _ := GetKubernetesDeployments("")
	ctx.JSON(http.StatusOK, d)
}

func GetAppHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	d, _ := GetKubernetesDeployment(name, "")
	ctx.JSON(http.StatusOK, d)
}

func DeleteAppHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	err := DeleteKubernetesDeployment(name, "")
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}

// Custom struct to represent the JSON data
type UpdateAppHandlerData struct {
	Name          string               `json:"name"`
	UpdateOptions metav1.UpdateOptions `json:"updateOptions"`
}

func UpdateAppHandler(ctx *gin.Context) {
	// Create an instance of UpdateAppHandlerData to bind the JSON data
	var data UpdateAppHandlerData

	// Bind the JSON data from the request body to the updateOptions struct
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	err := UpdateKubernetesDeployment(data.Name, "", data.UpdateOptions)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"updated": true})
}
