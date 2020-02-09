package model

type Page struct {
	Total   int         `json:"totalElements"`
	Size    int         `json:"payloadSize"`
	Payload interface{} `json:"payload"`
}

func NewPage(total, size int, payload interface{}) *Page {
	return &Page{
		Total:   total,
		Size:    size,
		Payload: payload,
	}
}
