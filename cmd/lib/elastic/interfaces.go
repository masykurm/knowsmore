package elastic

//go:generate moq -out client.go . Client
//go:generate moq -out search_service.go . SearchService

// Client wraps elastic.Client type
type Client interface{}

// SearchService wraps elastic.SearchService type
type SearchService interface{}
