package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type OrderRouter struct {
	OrderController OrderController
}

func NewRouter(db *sqlx.DB) OrderRouter {
	return OrderRouter{
		OrderController: &orderController{
			OrderUseCase: &orderUseCase{
				ProductRepo: repositories.NewProductRepository(db),
				OrderRepo:   repositories.NewOrderRepository(db),
			},
		},
	}
}

func (or *OrderRouter) Handle(router *fiber.App) {
	router.Get("/order", or.OrderController.GetAllOrders)
	router.Get("/order/:id", or.OrderController.GetOrderById)
	router.Post("/order", or.OrderController.CreateOrder)
	router.Put("/order/:id", or.OrderController.UpdateOrder)
	router.Delete("/order/:id", or.OrderController.DeleteOrder)
}
