package http

//go:generate moq -out request.go . Request
//go:generate moq -out response.go . Response
//go:generate moq -out client.go . Client

// Request wraps http.Request type
type Request interface{}

// Response wraps http.Response type
type Response interface{}

// Client wraps http.Client type
type Client interface{}
