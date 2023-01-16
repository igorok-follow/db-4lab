package models

type Detail struct {
	Name          string `db:"detail_name"`
	Weight        float32
	MaterialName  string `db:"material_name"`
	Amount        int32  `db:"details_amount"`
	MaterialsCost float32
}
