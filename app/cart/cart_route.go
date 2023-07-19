package cart

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type CartRouter struct {
	CartController CartController
}

func NewRouter(db *sqlx.DB) CartRouter {
	return CartRouter{
		CartController: &cartController{
			CartUseCase: &cartUseCase{
				CartRepo: repositories.NewCartRepository(db),
			},
		},
	}
}

func (cr *CartRouter) Handle(router *fiber.App) {
	router.Get("/cart/:userid", cr.CartController.GetCartByUserId)
	router.Put("/cart", cr.CartController.UpdateProductCart)
	router.Delete("/cart", cr.CartController.DeleteProductCart)
}
