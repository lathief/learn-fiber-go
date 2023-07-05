package category

import (
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type categoryUseCase struct {
	CategoryRepo repositories.CategoryRepository
}

type CategoryUseCase interface {
	GetCategoryById(id int) (CategoryDTO, error)
	CreateCategory(category CategoryDTO) error
}

func (cu *categoryUseCase) GetCategoryById(id int) (CategoryDTO, error) {
	//TODO implement me
	panic("implement me")
}
func (cu *categoryUseCase) CreateCategory(category CategoryDTO) error {
	var categorySave models.Category
	categorySave.Name = category.Name
	categorySave.Description = category.Description
	err := cu.CategoryRepo.Create(categorySave)
	if err != nil {
		return err
	}
	return nil
}
