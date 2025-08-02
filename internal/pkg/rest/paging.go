package rest

const (
	DefaultPageSize   = 10
	DefaultPageNumber = 0
)

type Paging struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

func NewPaging(pageSize, pageNumber int) *Paging {
	if pageSize <= 0 || pageSize > 50 {
		pageSize = DefaultPageSize
	}
	if pageNumber < 0 {
		pageNumber = DefaultPageNumber
	}
	return &Paging{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
}
