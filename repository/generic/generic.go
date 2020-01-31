package repository

type GenericRepository interface {
	Create(model interface{}) interface{}
	FetchOne(model interface{}, id int) interface{}
	Fetch(model interface{}) interface{}
	Update(model interface{}, id int) interface{}
	Delete(id int)
	Exists(id int) bool
}
