package repository

import (
	"context"
	"database-service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabaseRepository(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (d *Database) GetMaterials(ctx context.Context) ([]*models.Material, error) {
	var materials []*models.Material
	err := d.db.SelectContext(ctx, &materials, `
		select material_name, cost_per_gram from "public.materials" where active = '1';
	`)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func (d *Database) GetDetails(ctx context.Context) ([]*models.Detail, error) {
	var details []*models.Detail
	err := d.db.SelectContext(ctx, &details, `
		select detail_name, weight, material_name from "public.details" where active = '1';
	`)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func (d *Database) GetProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := d.db.SelectContext(ctx, &products, `
		select product_number, product_name from "public.products" where active = '1';
	`)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		var details []*models.Detail
		err = d.db.SelectContext(ctx, &details, `
			select detail_name, details_amount from "public.product_composition" where product_number = $1
		`, product.Id)
		if err != nil {
			return nil, err
		}

		product.Details = details
	}

	return products, nil
}

func (d *Database) DeleteMaterials(ctx context.Context, name string) error {
	_, err := d.db.ExecContext(ctx, `
		update "public.materials" set active = '0' where material_name = $1
	`, name)

	return err
}

func (d *Database) DeleteDetails(ctx context.Context, name string) error {
	_, err := d.db.ExecContext(ctx, `
		update "public.details" set active = '0' where detail_name = $1
	`, name)

	return err
}

func (d *Database) DeleteProducts(ctx context.Context, id int32) error {
	_, err := d.db.ExecContext(ctx, `
		update "public.products" set active = '0' where product_number = $1
	`, id)

	return err
}

func (d *Database) InsertMaterials(ctx context.Context, material *models.Material) error {
	_, err := d.db.ExecContext(ctx, `
		insert into "public.materials"(material_name, cost_per_gram) values($1, $2)
	`, material.Name, material.CostPerGram)

	return err
}

func (d *Database) InsertDetails(ctx context.Context, detail *models.Detail) error {
	_, err := d.db.ExecContext(ctx, `
		insert into "public.details"(detail_name, weight, material_name) values($1, $2, $3)
	`, detail.Name, detail.Weight, detail.MaterialName)

	return err
}

func (d *Database) InsertProducts(ctx context.Context, product *models.Product) (err error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	var lastId int
	err = tx.QueryRowContext(ctx, `
		insert into "public.products"(product_name) values($1) returning product_number
	`, product.Name).Scan(&lastId)
	if err != nil {
		log.Println(err)
	}

	for _, detail := range product.Details {
		_, err = tx.ExecContext(ctx, `
			insert into "public.product_composition"(product_number, detail_name, details_amount) values($1, $2, $3)
		`, lastId, detail.Name, detail.Amount)
		if err != nil {
			log.Println(err)
		}
	}

	if err != nil {
		err = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

func (d *Database) UpdateMaterials(ctx context.Context, material *models.Material, oldName string) error {
	_, err := d.db.ExecContext(ctx, `
		update "public.materials" set material_name = $1, cost_per_gram = $2 where material_name = $3
	`, material.Name, material.CostPerGram, oldName)

	return err
}

func (d *Database) UpdateDetails(ctx context.Context, detail *models.Detail, oldName string) error {
	_, err := d.db.ExecContext(ctx, `
		update "public.details" set detail_name = $1, weight = $2, material_name = $3 where detail_name = $4
	`, detail.Name, detail.Weight, detail.MaterialName, oldName)

	return err
}

func (d *Database) UpdateProducts(ctx context.Context, product *models.Product) (err error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = tx.ExecContext(ctx, `
		update "public.products" set product_name = $1 where product_number = $2
	`, product.Name, product.Id)
	if err != nil {
		log.Println(err)
	}

	_, err = tx.ExecContext(ctx, `
		delete from "public.product_composition" where product_number = $1
	`, product.Id)

	for _, detail := range product.Details {
		_, err = tx.ExecContext(ctx, `
			insert into "public.product_composition"(product_number, detail_name, details_amount) values($1, $2, $3)
		`, product.Id, detail.Name, detail.Amount)
		if err != nil {
			log.Println(err)
		}
	}

	if err != nil {
		err = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	return
}
