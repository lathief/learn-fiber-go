package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"strconv"
)

type orderController struct {
	OrderUseCase OrderUseCase
}

type OrderController interface {
	GetOrderById(ctx *fiber.Ctx) error
	CreateOrder(ctx *fiber.Ctx) error
	GetAllOrders(ctx *fiber.Ctx) error
	UpdateOrder(ctx *fiber.Ctx) error
	DeleteOrder(ctx *fiber.Ctx) error
}

func (oc orderController) GetOrderById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	res := oc.OrderUseCase.GetOrderById(s)
	return ctx.Status(res.Code).JSON(res)
}

func (oc orderController) CreateOrder(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (oc orderController) GetAllOrders(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (oc orderController) UpdateOrder(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (oc orderController) DeleteOrder(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
