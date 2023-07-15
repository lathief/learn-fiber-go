package category

import (
	"context"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type categoryUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}

type CategoryUseCase interface {
	GetCategoryById(ctx context.Context, id int) (categoryDTO dtos.CategoryDTO, err error)
	CreateCategory(ctx context.Context, category dtos.CategoryDTO) (err error)
	UpdateCategory(ctx context.Context, id int, category dtos.CategoryDTO) (err error)
	DeleteCategory(ctx context.Context, id int) (err error)
	GetAllCategories(ctx context.Context) (categoriesDTO []dtos.AllCategoryDTO, err error)
}

func (cu *categoryUseCase) GetCategoryById(ctx context.Context, id int) (categoryDTO dtos.CategoryDTO, err error) {
	getCategory, err := cu.CategoryRepo.GetById(ctx, int64(id))
	if err != nil {
		return dtos.CategoryDTO{}, err
	}
	categoryProducts, err := cu.ProductRepo.GetAllByCategoryId(ctx, getCategory.ID)
	if err != nil {
		return dtos.CategoryDTO{}, err
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
	categoryDTO = dtos.CategoryDTO{
		ID:          getCategory.ID,
		Name:        getCategory.Name,
		Description: getCategory.Description,
		Products:    products,
	}
	return categoryDTO, nil
}
func (cu *categoryUseCase) CreateCategory(ctx context.Context, category dtos.CategoryDTO) (err error) {
	var categorySave models.Category
	categorySave.Name = category.Name
	categorySave.Description = category.Description
	err = cu.CategoryRepo.Create(ctx, categorySave)
	return err
}
func (cu *categoryUseCase) UpdateCategory(ctx context.Context, id int, category dtos.CategoryDTO) (err error) {
	var categoryUpdate models.Category
	categoryUpdate.ID = int64(id)
	categoryUpdate.Name = category.Name
	categoryUpdate.Description = category.Description
	err = cu.CategoryRepo.Update(ctx, categoryUpdate)
	return err
}
func (cu *categoryUseCase) DeleteCategory(ctx context.Context, id int) (err error) {
	err = cu.CategoryRepo.Delete(ctx, int64(id))
	return err
}
func (cu *categoryUseCase) GetAllCategories(ctx context.Context) (categoriesDTO []dtos.AllCategoryDTO, err error) {
	categories, err := cu.CategoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range categories {
		categoriesDTO = append(categoriesDTO, dtos.AllCategoryDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return categoriesDTO, nil
}
