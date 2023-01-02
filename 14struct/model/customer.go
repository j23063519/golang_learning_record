package model

// exported struct
type Customer struct {
	// exported field
	Id int
	// exported field
	Name string
	// unexported field (only accessible inside package `model`)
	addr address
	// unexported field (only accessible inside package `model`)
	married bool
}
