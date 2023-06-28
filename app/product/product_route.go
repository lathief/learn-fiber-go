package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type ProductRouter struct {
	ProductController ProductController
}

func NewRouter(db *sqlx.DB) ProductRouter {
	return ProductRouter{
		ProductController: &productController{
			ProdcutUseCase: &productUseCase{
				ProductRepo: repositories.NewProductRepository(db),
			},
		},
	}
}

func (pr *ProductRouter) Handle(router *fiber.App) {
	router.Get("/product", pr.ProductController.GetAllProducts)
	router.Get("/product/:id", pr.ProductController.GetProductById)
	router.Post("/product", pr.ProductController.CreateProduct)
	router.Patch("/product/:id", pr.ProductController.UpdateProduct)
	router.Delete("/product/:id", pr.ProductController.DeleteProduct)
}
