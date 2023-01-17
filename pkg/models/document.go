package models

type Document1 struct {
	ProductName string  `db:"product_name"`
	DetailName  string  `db:"detail_name"`
	Cost        float32 `db:"cost"`
}

type Document2 struct {
	ProductName     string  `db:"product_name"`
	MaterialsWeight float32 `db:"mw"`
}
