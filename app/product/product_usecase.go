package product

import (
	"context"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type productUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}
type ProductUseCase interface {
	GetAllProducts(ctx context.Context) (productsDTO []dtos.AllProductsDTO, err error)
	GetProductById(ctx context.Context, id int) (productDTO dtos.ProductDTO, err error)
	CreateProduct(ctx context.Context, product dtos.ProductDTO) (err error)
	UpdateProduct(ctx context.Context, id int, product dtos.ProductDTO) (err error)
	DeleteProduct(ctx context.Context, id int) (err error)
}

func (pu *productUseCase) GetAllProducts(ctx context.Context) (productsDTO []dtos.AllProductsDTO, err error) {
	products, err := pu.ProductRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		productsDTO = append(productsDTO, dtos.AllProductsDTO{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CategoryId:  product.CategoryId,
		})
	}
	return productsDTO, nil
}
func (pu *productUseCase) GetProductById(ctx context.Context, id int) (productDTO dtos.ProductDTO, err error) {
	product, err := pu.ProductRepo.GetById(ctx, int64(id))
	if err != nil {
		return dtos.ProductDTO{}, err
	}
	categoryProduct, err := pu.CategoryRepo.GetById(ctx, product.CategoryId)
	if err != nil {
		return dtos.ProductDTO{}, err
	}
	productDTO = dtos.ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category: dtos.CategoryDTO{
			Name:        categoryProduct.Name,
			Description: categoryProduct.Description,
		},
	}
	return productDTO, nil
}
func (pu *productUseCase) CreateProduct(ctx context.Context, product dtos.ProductDTO) (err error) {
	var productSave models.Product
	productSave.Name = product.Name
	productSave.Price = product.Price
	productSave.Description = product.Description
	productSave.CategoryId = product.CategoryId

	err = pu.ProductRepo.Create(ctx, productSave)
	return err
}
func (pu *productUseCase) UpdateProduct(ctx context.Context, id int, product dtos.ProductDTO) (err error) {
	var productUpdate models.Product
	productUpdate.ID = int64(id)
	productUpdate.Name = product.Name
	productUpdate.Price = product.Price
	productUpdate.Description = product.Description
	productUpdate.CategoryId = product.CategoryId

	err = pu.ProductRepo.Update(ctx, productUpdate)
	return err
}
func (pu *productUseCase) DeleteProduct(ctx context.Context, id int) (err error) {
	err = pu.ProductRepo.Delete(ctx, int64(id))
	return err
}
