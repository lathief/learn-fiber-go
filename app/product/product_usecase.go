package product

import (
	"database/sql"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
	"strconv"
)

type productUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}
type ProductUseCase interface {
	GetAllProducts() handlers.ReturnResponse
	GetProductById(id int) handlers.ReturnResponse
	CreateProduct(product dtos.ProductDTO) handlers.ReturnResponse
	UpdateProduct(id int, product dtos.ProductDTO) handlers.ReturnResponse
	DeleteProduct(id int) handlers.ReturnResponse
}

func (pu *productUseCase) GetAllProducts() handlers.ReturnResponse {
	products, err := pu.ProductRepo.GetAll()
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var productsDTO []dtos.AllProductsDTO
	for _, product := range products {
		productsDTO = append(productsDTO, dtos.AllProductsDTO{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			CategoryId:  product.CategoryId,
		})
	}

	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    productsDTO,
	}
}
func (pu *productUseCase) GetProductById(id int) handlers.ReturnResponse {
	product, err := pu.ProductRepo.GetById(int64(id))
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	categoryProduct, err := pu.CategoryRepo.GetById(product.CategoryId)
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var productsDTO = dtos.ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category: dtos.CategoryDTO{
			Name:        categoryProduct.Name,
			Description: categoryProduct.Description,
		},
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    productsDTO,
	}
}
func (pu *productUseCase) CreateProduct(product dtos.ProductDTO) handlers.ReturnResponse {
	var productSave models.Product
	productSave.Name = product.Name
	productSave.Price = product.Price
	productSave.Description = product.Description
	productSave.CategoryId = product.CategoryId
	err := pu.ProductRepo.Create(productSave)
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
	}
}
func (pu *productUseCase) UpdateProduct(id int, product dtos.ProductDTO) handlers.ReturnResponse {
	var productUpdate models.Product
	productUpdate.ID = int64(id)
	productUpdate.Name = product.Name
	productUpdate.Price = product.Price
	productUpdate.Description = product.Description
	productUpdate.CategoryId = product.CategoryId
	err := pu.ProductRepo.Update(productUpdate)
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
	}
}
func (pu *productUseCase) DeleteProduct(id int) handlers.ReturnResponse {
	err := pu.ProductRepo.Delete(int64(id))
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
	}
}
