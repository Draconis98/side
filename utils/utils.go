package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Container struct {
	containerId string `json:"containerId"`
	core        string `json:"core"`
	memory      string `json:"memory"`
	status      string `json:"status"`
	createAt    string `json:"createAt"`
}

type ErrorMessage struct {
	message string `json:"message"`
}

type ContainerCreation struct {
	baseImage string `json:"baseImage"`
	core      string `json:"core"`
	memory    string `json:"memory"`
}

type ContainerExpansion struct {
	containerId string `json:"containerId"`
	newCore     string `json:"newCore"`
	newMemory   string `json:"newMemory"`
}

type ContainerDeletion struct {
	containerId string `json:"containerId"`
}

// ParseJson a generic function to parse json from request
func ParseJson[T any](request *http.Request) T {
	var data T
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	return data
}

// WriteJson a generic function to write json to response
func WriteJson[T any](response http.ResponseWriter, data T, statusCode int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	err := json.NewEncoder(response).Encode(data)
	if err != nil {
		log.Println(err)
	}
}
