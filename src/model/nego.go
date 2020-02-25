package model

//Nego is a generic interface that all NEGO models should implement
type Nego interface {

	//Valid uses validator package
	Valid() error
}
