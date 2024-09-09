package utils

type HeaderInfo struct {
	Username string `header:"X-Nickname" binding:"required"`
}

type Container struct {
	ContainerId string `json:"containerId"`
	Core        int    `json:"core"`
	Memory      int    `json:"memory"`
	Status      int    `json:"status"`
	CreateAt    string `json:"createAt"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ContainerCreation struct {
	BaseImage string `json:"baseImage"`
	Core      int    `json:"core"`
	Memory    int    `json:"memory"`
}

type ContainerExpansion struct {
	ContainerId string `json:"containerId"`
	NewCore     int    `json:"newCore"`
	NewMemory   int    `json:"newMemory"`
}

type ContainerDeletion struct {
	ContainerId string `json:"containerId"`
}

type ContainerRestore struct {
	ContainerId string `json:"containerId"`
	Core        int    `json:"core"`
	Memory      int    `json:"core"`
}
