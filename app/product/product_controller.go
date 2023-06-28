package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"strconv"
)

type productController struct {
	ProdcutUseCase ProductUseCase
}
type ProductController interface {
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

func (p *productController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := p.ProdcutUseCase.GetAllProducts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ResponseWithData{
		Code:    200,
		Message: "Success",
		Data:    products,
	})
}
func (p *productController) GetProductById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ResponseWithoutData{
			Code:    400,
			Message: "Bad Request",
		})
	}
	product, err := p.ProdcutUseCase.GetProductById(s)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ResponseWithData{
		Code:    200,
		Message: "Success",
		Data:    product,
	})
}

func (p *productController) CreateProduct(ctx *fiber.Ctx) error {
	var productReq ProductDTO
	if err := ctx.BodyParser(&productReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	err := p.ProdcutUseCase.CreateProduct(productReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ResponseWithoutData{
		Code:    200,
		Message: "Success",
	})
}

func (p *productController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ResponseWithoutData{
			Code:    400,
			Message: "Bad Request",
		})
	}
	var productReq ProductDTO
	if err = ctx.BodyParser(&productReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	err = p.ProdcutUseCase.UpdateProduct(id, productReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ResponseWithoutData{
		Code:    200,
		Message: "Success",
	})
}

func (p *productController) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ResponseWithoutData{
			Code:    400,
			Message: "Bad Request",
		})
	}
	err = p.ProdcutUseCase.DeleteProduct(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ResponseWithoutData{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ResponseWithoutData{
		Code:    200,
		Message: "Success",
	})
}
