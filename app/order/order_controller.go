package order

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/constant"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"net/http"
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

func (oc *orderController) GetAllOrders(ctx *fiber.Ctx) error {
	data, err := oc.OrderUseCase.GetAllOrders(ctx.Context())
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}
func (oc *orderController) GetOrderById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	data, err := oc.OrderUseCase.GetOrderById(ctx.Context(), s)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrOrderNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, constant.JoinMsgError(constant.ErrInternalServerError, err.Error()),
			http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}

func (oc *orderController) CreateOrder(ctx *fiber.Ctx) error {
	var orderReq dtos.OrderReqDTO
	if err := ctx.BodyParser(&orderReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := oc.OrderUseCase.CreateOrder(ctx.Context(), orderReq)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}

func (oc *orderController) UpdateOrder(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (oc *orderController) DeleteOrder(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
