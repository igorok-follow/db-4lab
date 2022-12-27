package repository

import (
	"context"
	"database-service/pkg/models"
	"github.com/jmoiron/sqlx"
)

type DatabaseRepository interface {
	GetMaterials(ctx context.Context) ([]*models.Material, error)
	GetDetails(ctx context.Context) ([]*models.Detail, error)
	GetProducts(ctx context.Context) ([]*models.Product, error)
	DeleteMaterials(ctx context.Context, name string) error
	DeleteDetails(ctx context.Context, name string) error
	DeleteProducts(ctx context.Context, id int32) error
	InsertMaterials(ctx context.Context, material *models.Material) error
	InsertDetails(ctx context.Context, detail *models.Detail) error
	InsertProducts(ctx context.Context, product *models.Product) error
	UpdateMaterials(ctx context.Context, material *models.Material) error
	UpdateDetails(ctx context.Context, detail *models.Detail) error
	UpdateProducts(ctx context.Context, product *models.Product) error
}

type Repositories struct {
	DatabaseRepository DatabaseRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		DatabaseRepository: NewDatabaseRepository(db),
	}
}
