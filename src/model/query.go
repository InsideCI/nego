package model

type QueryParams struct {
	Limit  int
	Offset int
	Page   int
	Order  []string
}
