package category

import (
	"database/sql"
	dtos2 "github.com/lathief/learn-fiber-go/app/dtos"
	"github.com/lathief/learn-fiber-go/app/models"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
	"strconv"
)

type categoryUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}

type CategoryUseCase interface {
	GetCategoryById(id int) (handlers.ReturnResponse, error)
	CreateCategory(category dtos2.CategoryDTO) handlers.ReturnResponse
	UpdateCategory(id int, category dtos2.CategoryDTO) handlers.ReturnResponse
	DeleteCategory(id int) handlers.ReturnResponse
	GetAllCategories() handlers.ReturnResponse
}

func (cu *categoryUseCase) GetCategoryById(id int) (handlers.ReturnResponse, error) {
	getCategory, err := cu.CategoryRepo.GetById(int64(id))
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}, err
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}, err
	}
	categoryProducts, err := cu.ProductRepo.GetAllByCategoryId(getCategory.ID)
	if err == sql.ErrNoRows {
		return handlers.ReturnResponse{
			Code:    404,
			Message: "Not Found: Data Not Found With id " + strconv.Itoa(id),
		}, err
	}
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}, err
	}
	var products []dtos2.ProductDTO
	for _, item := range categoryProducts {
		products = append(products, dtos2.ProductDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	var categoryDTO = dtos2.CategoryDTO{
		ID:          getCategory.ID,
		Name:        getCategory.Name,
		Description: getCategory.Description,
		Products:    products,
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    categoryDTO,
	}, nil
}
func (cu *categoryUseCase) CreateCategory(category dtos2.CategoryDTO) handlers.ReturnResponse {
	var categorySave models.Category
	categorySave.Name = category.Name
	categorySave.Description = category.Description
	err := cu.CategoryRepo.Create(categorySave)
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
func (cu *categoryUseCase) UpdateCategory(id int, category dtos2.CategoryDTO) handlers.ReturnResponse {
	var categoryUpdate models.Category
	categoryUpdate.ID = int64(id)
	categoryUpdate.Name = category.Name
	categoryUpdate.Description = category.Description
	err := cu.CategoryRepo.Update(categoryUpdate)
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
func (cu *categoryUseCase) DeleteCategory(id int) handlers.ReturnResponse {
	err := cu.CategoryRepo.Delete(int64(id))
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
func (cu *categoryUseCase) GetAllCategories() handlers.ReturnResponse {
	categories, err := cu.CategoryRepo.GetAll()
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var categoriesDTO []dtos2.AllCategoryDTO
	for _, item := range categories {
		categoriesDTO = append(categoriesDTO, dtos2.AllCategoryDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    categoriesDTO,
	}
}
