package service

import "github.com/Seunghoon-Oh/cloud-ml-pipeline-manager/data"

func GetPipelines() []string {
	return data.GetRedisData()
}

func CreatePipeline() string {
	// TODO: Create Kubernetes Object
	return data.CreateRedisData()
}
