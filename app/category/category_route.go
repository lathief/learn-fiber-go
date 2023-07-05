package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type CategoryRouter struct {
	CategoryController CategoryController
}

func NewRouter(db *sqlx.DB) CategoryRouter {
	return CategoryRouter{
		CategoryController: &categoryController{
			CategoryUseCase: &categoryUseCase{
				CategoryRepo: repositories.NewCategoryRepository(db),
			},
		},
	}
}

func (cr *CategoryRouter) Handle(router *fiber.App) {
	router.Get("/category/:id", cr.CategoryController.GetCategoryById)
	router.Post("/category", cr.CategoryController.CreateCategory)
}
