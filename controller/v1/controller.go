package v1

import (
	"net/http"

	"github.com/Seunghoon-Oh/cloud-ml-pipeline-manager/service"
	"github.com/gin-gonic/gin"
)

func GetPipelines(c *gin.Context) {
	data := service.GetPipelines()
	println("Response: " + data)
	c.String(http.StatusOK, data)
}
