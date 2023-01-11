package service

import (
	"context"
	"database-service/pkg/api"
	"database-service/pkg/models"
	"log"
)

type Database struct {
	deps *Dependencies
}

func NewDatabaseService(
	deps *Dependencies,
) *Database {
	return &Database{
		deps: deps,
	}
}

func (d *Database) GetMaterials(ctx context.Context, _ *api.Empty) (*api.GetMaterialsResponse, error) {
	materials, err := d.deps.Repository.DatabaseRepository.GetMaterials(ctx)
	if err != nil {
		return nil, err
	}

	materialsResp := make([]*api.Material, len(materials))
	for i := 0; i < len(materials); i++ {
		materialsResp[i] = &api.Material{
			Name:        materials[i].Name,
			CostPerGram: materials[i].CostPerGram,
		}
	}

	return &api.GetMaterialsResponse{
		Materials: materialsResp,
		Count:     int32(len(materialsResp)),
	}, nil
}

func (d *Database) GetDetails(ctx context.Context, _ *api.Empty) (*api.GetDetailsResponse, error) {
	details, err := d.deps.Repository.DatabaseRepository.GetDetails(ctx)
	if err != nil {
		return nil, err
	}

	detailsResp := make([]*api.Detail, len(details))
	for i := 0; i < len(details); i++ {
		detailsResp[i] = &api.Detail{
			Name:         details[i].Name,
			Weight:       details[i].Weight,
			MaterialName: details[i].MaterialName,
		}
	}

	return &api.GetDetailsResponse{
		Details: detailsResp,
		Count:   int32(len(detailsResp)),
	}, nil
}

func (d *Database) GetProducts(ctx context.Context, _ *api.Empty) (*api.GetProductsResponse, error) {
	products, err := d.deps.Repository.DatabaseRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	productsResp := make([]*api.Product, len(products))

	for i := 0; i < len(products); i++ {
		currentDetails := products[i].Details
		detailsResp := make([]*api.Detail, len(products[i].Details))
		for j := 0; j < len(currentDetails); j++ {
			detailsResp[j] = &api.Detail{
				Name:   currentDetails[j].Name,
				Amount: currentDetails[j].Amount,
			}

		}

		productsResp[i] = &api.Product{
			Id:      products[i].Id,
			Name:    products[i].Name,
			Details: detailsResp,
		}

		log.Println(productsResp[i].Details[0])
		log.Println(productsResp[i].Details[1])

	}

	return &api.GetProductsResponse{
		Products: productsResp,
		Count:    int32(len(productsResp)),
	}, nil
}
func (d *Database) DeleteMaterials(ctx context.Context, req *api.DeleteMaterialsReq) (*api.Empty, error) {
	err := d.deps.Repository.DatabaseRepository.DeleteMaterials(ctx, req.MaterialName)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) DeleteDetails(ctx context.Context, req *api.DeleteDetailsReq) (*api.Empty, error) {
	err := d.deps.Repository.DatabaseRepository.DeleteDetails(ctx, req.DetailName)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) DeleteProducts(ctx context.Context, req *api.DeleteProductsReq) (*api.Empty, error) {
	err := d.deps.Repository.DatabaseRepository.DeleteProducts(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) InsertMaterials(ctx context.Context, req *api.InsertMaterialsReq) (*api.Empty, error) {
	material := &models.Material{
		Name:        req.Material.Name,
		CostPerGram: req.Material.CostPerGram,
	}

	err := d.deps.Repository.DatabaseRepository.InsertMaterials(ctx, material)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) InsertDetails(ctx context.Context, req *api.InsertDetailsReq) (*api.Empty, error) {
	detail := &models.Detail{
		Name:         req.Detail.Name,
		Weight:       req.Detail.Weight,
		MaterialName: req.Detail.MaterialName,
	}

	err := d.deps.Repository.DatabaseRepository.InsertDetails(ctx, detail)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) InsertProducts(ctx context.Context, req *api.InsertProductsReq) (*api.Empty, error) {
	log.Println(req.Product.Name)
	log.Println(req.Product.Details)

	product := &models.Product{
		Id:      req.Product.Id,
		Name:    req.Product.Name,
		Details: nil,
	}

	details := make([]*models.Detail, len(req.Product.Details))
	for i := 0; i < len(req.Product.Details); i++ {
		details[i] = &models.Detail{
			Name:         req.Product.Details[i].Name,
			Weight:       req.Product.Details[i].Weight,
			MaterialName: req.Product.Details[i].MaterialName,
			Amount:       req.Product.Details[i].Amount,
		}
	}

	product.Details = details

	err := d.deps.Repository.DatabaseRepository.InsertProducts(ctx, product)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) UpdateMaterials(ctx context.Context, req *api.UpdateMaterialsReq) (*api.Empty, error) {
	material := &models.Material{
		Name:        req.Material.Name,
		CostPerGram: req.Material.CostPerGram,
	}

	err := d.deps.Repository.DatabaseRepository.UpdateMaterials(ctx, material, req.Material.OldName)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) UpdateDetails(ctx context.Context, req *api.UpdateDetailsReq) (*api.Empty, error) {
	detail := &models.Detail{
		Name:         req.Detail.Name,
		Weight:       req.Detail.Weight,
		MaterialName: req.Detail.MaterialName,
	}

	err := d.deps.Repository.DatabaseRepository.UpdateDetails(ctx, detail, req.Detail.OldName)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (d *Database) UpdateProducts(ctx context.Context, req *api.UpdateProductsReq) (*api.Empty, error) {
	product := &models.Product{
		Id:      req.Product.Id,
		Name:    req.Product.Name,
		Details: nil,
	}

	details := make([]*models.Detail, len(req.Product.Details))
	for i := 0; i < len(req.Product.Details); i++ {
		details[i] = &models.Detail{
			Name:         req.Product.Details[i].Name,
			Weight:       req.Product.Details[i].Weight,
			MaterialName: req.Product.Details[i].MaterialName,
			Amount:       req.Product.Details[i].Amount,
		}
	}

	product.Details = details

	err := d.deps.Repository.DatabaseRepository.UpdateProducts(ctx, product)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
