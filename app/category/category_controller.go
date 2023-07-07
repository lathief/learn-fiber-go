package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"strconv"
)

type categoryController struct {
	CategoryUseCase CategoryUseCase
}
type CategoryController interface {
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	GetAllCategories(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

func (cc *categoryController) GetCategoryById(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	Category, err := cc.CategoryUseCase.GetCategoryById(s)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    Category,
	})
}
func (cc *categoryController) CreateCategory(ctx *fiber.Ctx) error {
	var categoryReq dtos.CategoryDTO
	if err := ctx.BodyParser(&categoryReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	res := cc.CategoryUseCase.CreateCategory(categoryReq)
	return ctx.Status(res.Code).JSON(res)
}
func (cc *categoryController) GetAllCategories(ctx *fiber.Ctx) error {
	res := cc.CategoryUseCase.GetAllCategories()
	return ctx.Status(res.Code).JSON(res)
}
func (cc *categoryController) UpdateCategory(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	var categoryReq dtos.CategoryDTO
	if err = ctx.BodyParser(&categoryReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		})
	}
	res := cc.CategoryUseCase.UpdateCategory(s, categoryReq)
	return ctx.Status(res.Code).JSON(res)
}
func (cc *categoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(handlers.ReturnResponse{
			Code:    400,
			Message: "Bad Request",
		})
	}
	res := cc.CategoryUseCase.DeleteCategory(id)
	return ctx.Status(res.Code).JSON(res)
}
