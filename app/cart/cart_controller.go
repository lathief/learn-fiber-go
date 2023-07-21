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
	userId := ctx.Locals("userId").(string)
	id, err := strconv.Atoi(userId)
	if err != nil {
		return handlers.HandleResponse(ctx, constant.ErrInternalServerError.Error(), http.StatusInternalServerError)
	}
	data, err := cc.CartUseCase.GetCartByUserId(ctx.Context(), id)
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
	if err := ctx.BodyParser(&items); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	userId := ctx.Locals("userId").(string)
	id, err := strconv.Atoi(userId)
	if err != nil {
		return handlers.HandleResponse(ctx, constant.ErrInternalServerError.Error(), http.StatusInternalServerError)
	}
	err = cc.CartUseCase.UpdateProductInCart(ctx.Context(), id, items)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}

func (cc *cartController) DeleteProductCart(ctx *fiber.Ctx) error {
	var items dtos.CartProductIdDTO
	if err := ctx.BodyParser(&items); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	userId := ctx.Locals("userId").(string)
	id, err := strconv.Atoi(userId)
	if err != nil {
		return handlers.HandleResponse(ctx, constant.ErrInternalServerError.Error(), http.StatusInternalServerError)
	}
	err = cc.CartUseCase.DeleteProductsInCart(ctx.Context(), id, items)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
