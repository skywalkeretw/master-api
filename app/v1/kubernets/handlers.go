package kubernets

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	d, _ := GetKubernetesDeployments()
	ctx.JSON(http.StatusOK, d)
}

func GetAppHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	d, _ := GetKubernetesDeployment(name, "")
	ctx.JSON(http.StatusOK, d)
}

func DeleteAppHandler(ctx *gin.Context) {
	d, _ := GetKubernetesDeployments()
	ctx.JSON(http.StatusOK, d)
}

func UpdateAppHandler(ctx *gin.Context) {
	d, _ := GetKubernetesDeployments()
	ctx.JSON(http.StatusOK, d)
}

func DeploymentHandler(ctx *gin.Context) {
	d, _ := GetKubernetesDeployments()
	ctx.JSON(http.StatusOK, d)
}
