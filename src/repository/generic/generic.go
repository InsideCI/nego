package generic

type GenericRepository interface {
	Create(model interface{}) interface{}
	Count(model interface{}) int
	Delete(id int)
	FetchOne(model interface{}, id int) interface{}
	Fetch(model interface{}) interface{}
	Update(model interface{}, id int) interface{}
	Exists(id int) bool
}
