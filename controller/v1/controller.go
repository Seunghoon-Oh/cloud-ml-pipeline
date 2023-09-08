package v1

import (
	"net/http"

	"github.com/Seunghoon-Oh/cloud-ml-pipeline-manager/service"
	"github.com/gin-gonic/gin"
)

func GetPipelines(c *gin.Context) {
	data := service.GetPipelines()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreatePipeline(c *gin.Context) {
	data := service.CreatePipeline()
	c.String(http.StatusOK, data)
}
