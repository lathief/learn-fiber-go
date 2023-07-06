package category

import (
	"database/sql"
	"github.com/lathief/learn-fiber-go/app/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
	"strconv"
)

type categoryUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}

type CategoryUseCase interface {
	GetCategoryById(id int) (handlers.GetResponse, error)
	CreateCategory(category dtos.CategoryDTO) handlers.GetResponse
	UpdateCategory(id int, category dtos.CategoryDTO) handlers.GetResponse
	DeleteCategory(id int) handlers.GetResponse
	GetAllCategories() handlers.GetResponse
}

func (cu *categoryUseCase) GetCategoryById(id int) (handlers.GetResponse, error) {
	getCategory, err := cu.CategoryRepo.GetById(int64(id))
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
	categoryProducts, err := cu.ProductRepo.GetAllByCategoryId(getCategory.ID)
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
	var products []dtos.ProductDTO
	for _, item := range categoryProducts {
		products = append(products, dtos.ProductDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	var categoryDTO = dtos.CategoryDTO{
		ID:          getCategory.ID,
		Name:        getCategory.Name,
		Description: getCategory.Description,
		Products:    products,
	}
	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
		Data:    categoryDTO,
	}, nil
}
func (cu *categoryUseCase) CreateCategory(category dtos.CategoryDTO) handlers.GetResponse {
	var categorySave models.Category
	categorySave.Name = category.Name
	categorySave.Description = category.Description
	err := cu.CategoryRepo.Create(categorySave)
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
func (cu *categoryUseCase) UpdateCategory(id int, category dtos.CategoryDTO) handlers.GetResponse {
	var categoryUpdate models.Category
	categoryUpdate.ID = int64(id)
	categoryUpdate.Name = category.Name
	categoryUpdate.Description = category.Description
	err := cu.CategoryRepo.Update(categoryUpdate)
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
func (cu *categoryUseCase) DeleteCategory(id int) handlers.GetResponse {
	err := cu.CategoryRepo.Delete(int64(id))
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
func (cu *categoryUseCase) GetAllCategories() handlers.GetResponse {
	categories, err := cu.CategoryRepo.GetAll()
	if err != nil {
		return handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var categoriesDTO []dtos.AllCategoryDTO
	for _, item := range categories {
		categoriesDTO = append(categoriesDTO, dtos.AllCategoryDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return handlers.GetResponse{
		Code:    200,
		Message: "Success",
		Data:    categoriesDTO,
	}
}
