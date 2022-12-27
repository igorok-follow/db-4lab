package models

type Product struct {
	Id      int32
	Name    string
	Details []*Detail
}
