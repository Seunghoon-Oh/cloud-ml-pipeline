package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-pipeline-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-pipeline-manager/data"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/pipelines", v1.GetPipelines)
	r.POST("/pipeline", v1.CreatePipeline)

	return r
}

func main() {
	data.SetupRedisClient()
	r := setupRouter()
	r.Run(":8082")
}
