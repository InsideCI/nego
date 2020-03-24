package models

type Page struct {
	Total       int         `json:"totalElements"`
	TotalPages  int         `json:"totalPages"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PayloadSize int         `json:"payloadSize"`
	Payload     interface{} `json:"payload"`
}

func NewPage(total, limit, page, totalPages, payloadSize int, payload interface{}) *Page {
	return &Page{
		Total:       total,
		Limit:       limit,
		Page:        page,
		TotalPages:  totalPages,
		PayloadSize: payloadSize,
		Payload:     payload,
	}
}
