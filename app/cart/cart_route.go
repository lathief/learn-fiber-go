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
				ProductRepo:  repositories.NewProductRepository(db),
				CategoryRepo: repositories.NewCategoryRepository(db),
			},
		},
	}
}

func (cr *CartRouter) Handle(router *fiber.App) {
	router.Get("/cart/:id", cr.CartController.GetCartById)
	router.Post("/cart", cr.CartController.CreateCart)
	router.Get("/cart", cr.CartController.GetAllCarts)
	router.Put("/cart/:id", cr.CartController.UpdateCart)
	router.Delete("/cart/:id", cr.CartController.DeleteCart)
}
