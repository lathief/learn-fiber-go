package product

import (
	"database/sql"
	"github.com/lathief/learn-fiber-go/app/category"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
	"strconv"
)

type productUseCase struct {
	ProductRepo repositories.ProductRepository
}
type ProductUseCase interface {
	GetAllProducts() (handlers.GetResponse, error)
	GetProductById(id int) (handlers.GetResponse, error)
	CreateProduct(product ProductDTO) handlers.GetResponse
	UpdateProduct(id int, product ProductDTO) handlers.GetResponse
	DeleteProduct(id int) handlers.GetResponse
}

func (p *productUseCase) GetAllProducts() (handlers.GetResponse, error) {
	products, err := p.ProductRepo.GetAll()
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}, err
	}
	var productsDTO []AllProductsDTO
	for _, product := range products {
		productsDTO = append(productsDTO, AllProductsDTO{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CategoryId:  product.CategoryId,
		})
	}

	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
		Data:    productsDTO,
	}, nil
}
func (p *productUseCase) GetProductById(id int) (handlers.GetResponse, error) {
	product, err := p.ProductRepo.GetById(int64(id))
	if err == sql.ErrNoRows {
		return handlers.GetResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}, err
	}
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}, err
	}
	var productsDTO = ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category: category.CategoryDTO{
			Name:        product.Category.Name,
			Description: product.Category.Description,
		},
	}
	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
		Data:    productsDTO,
	}, nil
}
func (p *productUseCase) CreateProduct(product ProductDTO) handlers.GetResponse {
	var productSave models.Product
	productSave.Name = product.Name
	productSave.Price = product.Price
	productSave.Description = product.Description
	productSave.CategoryId = product.CategoryId
	err := p.ProductRepo.Create(productSave)
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
	}
}
func (p *productUseCase) UpdateProduct(id int, product ProductDTO) handlers.GetResponse {
	var productUpdate models.Product
	productUpdate.ID = int64(id)
	productUpdate.Name = product.Name
	productUpdate.Price = product.Price
	productUpdate.Description = product.Description
	productUpdate.CategoryId = product.CategoryId
	err := p.ProductRepo.Update(productUpdate)
	if err == sql.ErrNoRows {
		return handlers.GetResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
	}
}
func (p *productUseCase) DeleteProduct(id int) handlers.GetResponse {
	err := p.ProductRepo.Delete(int64(id))
	if err == sql.ErrNoRows {
		return handlers.GetResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
	}
}
