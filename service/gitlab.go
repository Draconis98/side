package service

import (
	"fmt"
	gitlab "github.com/xanzy/go-gitlab"
)

func TriggerPipeline(node, containerName string) (bool, int) {
	git, err := gitlab.NewClient("73UN5BsDyQXvsNUBW5qd", gitlab.WithBaseURL("https://gitlab.agileserve.org.cn:8001/api/v4"))
	if err != nil {
		fmt.Println("Error creating gitlab client:", err)
		return false, 0
	}

	projectID := "15101"
	refName := "trigger"
	Node := "NODE"
	Name := "containerName"

	pipeline, _, err := git.Pipelines.CreatePipeline(projectID, &gitlab.CreatePipelineOptions{
		Ref: &refName,
		Variables: &[]*gitlab.PipelineVariableOptions{
			{
				Key:   &Node,
				Value: &node,
			},
			{
				Key:   &Name,
				Value: &containerName,
			},
		},
	})
	if err != nil {
		fmt.Println("Error triggering pipeline:", err)
		return false, 0
	}

	fmt.Println("Pipeline triggered:", pipeline.ID, refName)

	return true, pipeline.ID
}

func CheckPipelineStatus(pipelineID int) bool {
	git, err := gitlab.NewClient("73UN5BsDyQXvsNUBW5qd", gitlab.WithBaseURL("https://gitlab.agileserve.org.cn:8001/api/v4"))
	if err != nil {
		fmt.Println("Error creating gitlab client:", err)
		return false
	}

	status := "created"
	for {
		pipeline, _, err := git.Pipelines.GetPipeline("15101", pipelineID)
		if err != nil {
			fmt.Println("Error getting pipeline:", err)
			return false
		}

		if pipeline.Status == "created" || pipeline.Status == "pending" || pipeline.Status == "running" {
			continue
		}

		status = pipeline.Status
		break
	}

	switch status {
	case "success":
		fmt.Println("Pipeline success")
		return true
	case "failed":
		fmt.Println("Pipeline failed")
		return false
	case "canceled":
		fmt.Println("Pipeline canceled")
		return false
	}

	fmt.Println("Pipeline status unknown")

	return false
}
