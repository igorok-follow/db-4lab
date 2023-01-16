package models

type Product struct {
	Id              int32  `db:"product_number"`
	Name            string `db:"product_name"`
	MaterialsWeight float32
	Details         []*Detail
}
