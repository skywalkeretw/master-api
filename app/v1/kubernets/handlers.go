package kubernets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "k8s.io/api/apps/v1"
	_ "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Summary Get Kubernetes Pods
// @Description Get a list of all Kubernetes Pods
// @Tags Pods
// @Accept json
// @Produce json
// @Success 200 {array} v1.Pod
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
// @Success 200 {array} v1.Deployment
// @Router /api/v1/app [get]
func GetAllAppsHandler(ctx *gin.Context) {
	d, err := GetKubernetesDeployments("")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, d)
}

// @Summary Get a Kubernetes Deployment
// @Description Get information about a Kubernetes Deployment by name
// @Tags Deployments
// @Accept json
// @Produce json
// @Param name path string true "Deployment Name"  example: my-deployment
// @Success 200 {object} v1.Deployment
// @Router /api/v1/apps/{name} [get]
func GetAppHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	d, _ := GetKubernetesDeployment(name, "")
	ctx.JSON(http.StatusOK, d)
}

// @Summary Delete a Kubernetes Deployment
// @Description Delete a Kubernetes Deployment by name
// @Tags Deployments
// @Accept json
// @Produce json
// @Param name path string true "Deployment Name" example: my-deployment
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/apps/{name} [delete]
func DeleteAppHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	err := DeleteKubernetesDeployment(name, "")
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}

// @Summary Update Request Data
// @Description Represents data for updating a Kubernetes Deployment
type UpdateAppHandlerData struct {
	Name          string               `json:"name"`
	UpdateOptions metav1.UpdateOptions `json:"updateOptions"`
}

// @Summary Update a Kubernetes Deployment
// @Description Update a Kubernetes Deployment by name with specified options
// @Tags Deployments
// @Accept json
// @Produce json
// @Param data body UpdateAppHandlerData true "Update Request Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/v1/apps [put]
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
