package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"strconv"
)

type categoryController struct {
	CategoryUseCase CategoryUseCase
}
type CategoryController interface {
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
}

func (cc *categoryController) GetCategoryById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.GetResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	Category, err := cc.CategoryUseCase.GetCategoryById(s)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.GetResponse{
		Code:    200,
		Message: "Success",
		Data:    Category,
	})
}
func (cc *categoryController) CreateCategory(ctx *fiber.Ctx) error {
	var categoryReq CategoryDTO
	if err := ctx.BodyParser(&categoryReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	err := cc.CategoryUseCase.CreateCategory(categoryReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.GetResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.GetResponse{
		Code:    200,
		Message: "Success",
	})
}
