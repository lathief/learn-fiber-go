package product

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

type productController struct {
	ProductUseCase ProductUseCase
}
type ProductController interface {
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

func (pc *productController) GetAllProducts(ctx *fiber.Ctx) error {
	data, err := pc.ProductUseCase.GetAllProducts(ctx.Context())
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}
func (pc *productController) GetProductById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	data, err := pc.ProductUseCase.GetProductById(ctx.Context(), s)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrProductNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}
func (pc *productController) CreateProduct(ctx *fiber.Ctx) error {
	var productReq dtos.ProductDTO
	if err := ctx.BodyParser(&productReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := pc.ProductUseCase.CreateProduct(ctx.Context(), productReq)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
func (pc *productController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	var productReq dtos.ProductDTO
	if err = ctx.BodyParser(&productReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err = pc.ProductUseCase.UpdateProduct(ctx.Context(), id, productReq)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrProductNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
func (pc *productController) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = pc.ProductUseCase.DeleteProduct(ctx.Context(), id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrProductNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
