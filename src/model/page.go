package model

type Page struct {
	Total      int         `json:"totalElements"`
	Offset     int         `json:"offset"`
	Page       int         `json:"page"`
	TotalPages int         `json:"totalPages"`
	Payload    interface{} `json:"payload"`
}

func NewPage(total, offset, page, totalPages int, payload interface{}) *Page {
	return &Page{
		Total:      total,
		Offset:     offset,
		Page:       page,
		TotalPages: totalPages,
		Payload:    payload,
	}
}
