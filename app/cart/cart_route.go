package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/middleware"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type CartRouter struct {
	Security       middleware.SecurityInterface
	CartController CartController
}

func NewRouter(db *sqlx.DB) CartRouter {
	return CartRouter{
		Security: middleware.NewSecurityRepo(db),
		CartController: &cartController{
			CartUseCase: &cartUseCase{
				ProductRepo: repositories.NewProductRepository(db),
				CartRepo:    repositories.NewCartRepository(db),
			},
		},
	}
}

func (cr *CartRouter) Handle(router *fiber.App) {
	cartRouter := router.Group("/cart")
	cartRouter.Use(cr.Security.Authentication)
	cartRouter.Get("/", cr.CartController.GetCartByUserId)
	cartRouter.Put("/", cr.CartController.UpdateProductCart)
	cartRouter.Delete("/", cr.CartController.DeleteProductCart)
}
