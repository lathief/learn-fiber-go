package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/middleware"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type ProductRouter struct {
	Security          middleware.SecurityInterface
	ProductController ProductController
}

func NewRouter(db *sqlx.DB) ProductRouter {
	return ProductRouter{
		Security: middleware.NewSecurityRepo(db),
		ProductController: &productController{
			ProductUseCase: &productUseCase{
				ProductRepo:  repositories.NewProductRepository(db),
				CategoryRepo: repositories.NewCategoryRepository(db),
			},
		},
	}
}

func (pr *ProductRouter) Handle(router *fiber.App) {
	router.Get("/product", pr.ProductController.GetAllProducts)
	router.Get("/product/:id", pr.ProductController.GetProductById)
	router.Post("/product", pr.Security.Authentication, pr.ProductController.CreateProduct)
	router.Put("/product/:id", pr.ProductController.UpdateProduct)
	router.Delete("/product/:id", pr.ProductController.DeleteProduct)
}
