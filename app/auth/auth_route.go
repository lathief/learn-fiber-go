package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthRouter struct {
	AuthController AuthController
}

func NewRouter(db *sqlx.DB) AuthRouter {
	return AuthRouter{
		AuthController: &authController{
			AuthUseCase: &authUseCase{},
		},
	}
}

func (pr *AuthRouter) Handle(router *fiber.App) {
	router.Post("/login", pr.AuthController.Login)
	router.Post("/register", pr.AuthController.Register)
	router.Post("/whoami", pr.AuthController.Whoami)
}
