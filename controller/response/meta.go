package response

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type LinkResponse struct {
	Meta
	URL string `json:"url"`
}

type Response struct {
	Meta
	Data interface{} `json:"data"`
}
