package category

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
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	data, err := cc.CategoryUseCase.GetCategoryById(ctx.Context(), s)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCategoryNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}
func (cc *categoryController) CreateCategory(ctx *fiber.Ctx) error {
	var categoryReq dtos.CategoryDTO
	if err := ctx.BodyParser(&categoryReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := cc.CategoryUseCase.CreateCategory(ctx.Context(), categoryReq)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
func (cc *categoryController) GetAllCategories(ctx *fiber.Ctx) error {
	data, err := cc.CategoryUseCase.GetAllCategories(ctx.Context())
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, data, "Success", http.StatusOK)
}
func (cc *categoryController) UpdateCategory(ctx *fiber.Ctx) error {
	s, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	var categoryReq dtos.CategoryDTO
	if err = ctx.BodyParser(&categoryReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err = cc.CategoryUseCase.UpdateCategory(ctx.Context(), s, categoryReq)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCategoryNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
func (cc *categoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusBadRequest)
	}
	err = cc.CategoryUseCase.DeleteCategory(ctx.Context(), id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return handlers.HandleResponse(ctx, constant.ErrCategoryNotFound.Error(), http.StatusNotFound)
	}
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "Success", http.StatusOK)
}
