package member

type (
	// Member : describe member struct
	Member struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// Request : describe member struct for request
	Request struct {
		Keyword string `json:"keyword"`
	}

	// Response : describe member struct for response
	Response struct {
		Members []Member `json:"members"`
		Status  int      `json:"status"`
		Error   string   `json:"error"`
	}
)
