package cart

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

type cartController struct {
	CartUseCase CartUseCase
}
type CartController interface {
	GetCartByUserId(ctx *fiber.Ctx) error
	UpdateProductCart(ctx *fiber.Ctx) error
	DeleteProductCart(ctx *fiber.Ctx) error
}

func (cc *cartController) GetCartByUserId(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	data, err := cc.CartUseCase.GetCartByUserId(ctx.Context(), s)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCartNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}

func (cc *cartController) UpdateProductCart(ctx *fiber.Ctx) error {
	var items dtos.Item
	var userId int
	if err := ctx.BodyParser(&items); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := cc.CartUseCase.UpdateProductInCart(ctx.Context(), userId, items)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCategoryNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}

func (cc *cartController) DeleteProductCart(ctx *fiber.Ctx) error {
	var items dtos.CartProductIdDTO
	var userId int
	if err := ctx.BodyParser(&items); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := cc.CartUseCase.DeleteProductsInCart(ctx.Context(), userId, items)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCategoryNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
