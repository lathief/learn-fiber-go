package product

import (
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type productUseCase struct {
	ProductRepo repositories.ProductRepository
}
type ProductUseCase interface {
	GetAllProducts() ([]ProductDTO, error)
	GetProductById(id int) (ProductDTO, error)
	CreateProduct(product ProductDTO) error
	UpdateProduct(id int, product ProductDTO) error
	DeleteProduct(id int) error
}

func (p *productUseCase) GetAllProducts() ([]ProductDTO, error) {
	products, err := p.ProductRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var productsDTO []ProductDTO
	for _, product := range products {
		productsDTO = append(productsDTO, ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	return productsDTO, nil
}
func (p *productUseCase) GetProductById(id int) (ProductDTO, error) {
	product, err := p.ProductRepo.GetById(int64(id))
	if err != nil {
		return ProductDTO{}, err
	}
	var productsDTO = ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return productsDTO, nil
}
func (p *productUseCase) CreateProduct(product ProductDTO) error {
	var productSave models.Product
	productSave.Name = product.Name
	productSave.Price = product.Price
	productSave.Description = product.Description
	err := p.ProductRepo.Create(productSave)
	if err != nil {
		return err
	}
	return nil
}
func (p *productUseCase) UpdateProduct(id int, product ProductDTO) error {
	var productUpdate models.Product
	productUpdate.ID = int64(id)
	productUpdate.Name = product.Name
	productUpdate.Price = product.Price
	productUpdate.Description = product.Description
	err := p.ProductRepo.Update(productUpdate)
	if err != nil {
		return err
	}
	return nil
}
func (p *productUseCase) DeleteProduct(id int) error {
	err := p.ProductRepo.Delete(int64(id))
	if err != nil {
		return err
	}
	return nil
}
