package repository

import (
	"context"
	"database-service/pkg/models"
	"github.com/jmoiron/sqlx"
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
		select * from materials;
	`)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func (d *Database) GetDetails(ctx context.Context) ([]*models.Detail, error) {
	var details []*models.Detail
	err := d.db.SelectContext(ctx, &details, `
		select * from details;
	`)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func (d *Database) GetProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := d.db.SelectContext(ctx, &products, `
		select * from products;
	`)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d *Database) DeleteMaterials(ctx context.Context, name string) error {
	return nil
}
func (d *Database) DeleteDetails(ctx context.Context, name string) error {
	return nil
}
func (d *Database) DeleteProducts(ctx context.Context, id int32) error {
	return nil
}
func (d *Database) InsertMaterials(ctx context.Context, material *models.Material) error {
	return nil
}
func (d *Database) InsertDetails(ctx context.Context, detail *models.Detail) error {
	return nil
}
func (d *Database) InsertProducts(ctx context.Context, product *models.Product) error {
	return nil
}
func (d *Database) UpdateMaterials(ctx context.Context, material *models.Material) error {
	return nil
}
func (d *Database) UpdateDetails(ctx context.Context, detail *models.Detail) error {
	return nil
}
func (d *Database) UpdateProducts(ctx context.Context, product *models.Product) error {
	return nil
}
