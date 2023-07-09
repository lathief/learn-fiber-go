package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/app/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
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
	res := pc.ProductUseCase.GetAllProducts()
	return ctx.Status(res.Code).JSON(res)
}
func (pc *productController) GetProductById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	res := pc.ProductUseCase.GetProductById(s)
	return ctx.Status(res.Code).JSON(res)
}
func (pc *productController) CreateProduct(ctx *fiber.Ctx) error {
	var productReq dtos.ProductDTO
	if err := ctx.BodyParser(&productReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	res := pc.ProductUseCase.CreateProduct(productReq)
	return ctx.Status(res.Code).JSON(res)
}
func (pc *productController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	var productReq dtos.ProductDTO
	if err = ctx.BodyParser(&productReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	res := pc.ProductUseCase.UpdateProduct(id, productReq)
	return ctx.Status(res.Code).JSON(res)
}
func (pc *productController) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	res := pc.ProductUseCase.DeleteProduct(id)
	return ctx.Status(res.Code).JSON(res)
}
