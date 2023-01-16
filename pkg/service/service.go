package service

import (
	"context"

	"database-service/pkg/api"
	"database-service/pkg/repository"
)

type DatabaseService interface {
	GetMaterials(ctx context.Context, req *api.Empty) (*api.GetMaterialsResponse, error)
	GetDetails(ctx context.Context, req *api.Empty) (*api.GetDetailsResponse, error)
	GetProducts(ctx context.Context, req *api.Empty) (*api.GetProductsResponse, error)
	DeleteMaterials(ctx context.Context, req *api.DeleteMaterialsReq) (*api.Empty, error)
	DeleteDetails(ctx context.Context, req *api.DeleteDetailsReq) (*api.Empty, error)
	DeleteProducts(ctx context.Context, req *api.DeleteProductsReq) (*api.Empty, error)
	InsertMaterials(ctx context.Context, req *api.InsertMaterialsReq) (*api.Empty, error)
	InsertDetails(ctx context.Context, req *api.InsertDetailsReq) (*api.Empty, error)
	InsertProducts(ctx context.Context, req *api.InsertProductsReq) (*api.Empty, error)
	UpdateMaterials(ctx context.Context, req *api.UpdateMaterialsReq) (*api.Empty, error)
	UpdateDetails(ctx context.Context, req *api.UpdateDetailsReq) (*api.Empty, error)
	UpdateProducts(ctx context.Context, req *api.UpdateProductsReq) (*api.Empty, error)
	Document1(ctx context.Context, req *api.GetDocument1Req) (*api.GetDocument1Resp, error)
	Document2(ctx context.Context, req *api.GetDocument2Req) (*api.GetDocument2Resp, error)
}

type Services struct {
	DatabaseService DatabaseService
}

type Dependencies struct {
	Repository *repository.Repositories
}

func NewServices(deps *Dependencies) *Services {
	databaseService := NewDatabaseService(deps)

	return &Services{
		DatabaseService: databaseService,
	}
}
