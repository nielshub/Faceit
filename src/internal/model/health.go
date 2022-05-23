package model

type HealthResponse struct {
	Message string
	Status  int
	Version string
	Stack   string
}
