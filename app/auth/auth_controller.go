package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"net/http"
)

type authController struct {
	AuthUseCase AuthUseCase
}

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Whoami(ctx *fiber.Ctx) error
	CreateAccessToken(ctx *fiber.Ctx) error
}

func (a *authController) Login(ctx *fiber.Ctx) error {
	var loginReq dtos.LoginDTO
	if err := ctx.BodyParser(&loginReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	resp, err := a.AuthUseCase.Login(ctx.Context(), loginReq)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, resp, "login success", http.StatusOK)
}

func (a *authController) Register(ctx *fiber.Ctx) error {
	var registerReq dtos.RegisterDTO
	if err := ctx.BodyParser(&registerReq); err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	err := a.AuthUseCase.Register(ctx.Context(), registerReq)
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponse(ctx, "register success", http.StatusOK)
}

func (a *authController) Whoami(ctx *fiber.Ctx) error {
	user, err := a.AuthUseCase.Whoami(ctx.Context(), ctx.Locals("userId").(string))
	if err != nil {
		return handlers.HandleResponse(ctx, err.Error(), http.StatusInternalServerError)
	}
	return handlers.HandleResponseWithData(ctx, user, "success", http.StatusOK)
}
func (a *authController) CreateAccessToken(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
