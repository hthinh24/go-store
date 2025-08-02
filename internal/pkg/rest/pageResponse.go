package rest

type PageResponse struct {
	PageSize   int         `json:"page_size"`
	PageNumber int         `json:"page_number"`
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
