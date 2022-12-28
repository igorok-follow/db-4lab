package models

type Material struct {
	Name        string  `db:"material_name"`
	CostPerGram float32 `db:"cost_per_gram"`
}
