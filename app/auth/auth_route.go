package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/middleware"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type AuthRouter struct {
	Security       middleware.SecurityInterface
	AuthController AuthController
}

func NewRouter(db *sqlx.DB) AuthRouter {
	return AuthRouter{
		Security: middleware.NewSecurityRepo(db),
		AuthController: &authController{
			AuthUseCase: &authUseCase{
				UserRepo: repositories.NewUserRepository(db),
				RoleRepo: repositories.NewRoleRepository(db),
				CartRepo: repositories.NewCartRepository(db),
			},
		},
	}
}

func (pr *AuthRouter) Handle(router *fiber.App) {
	router.Post("/login", pr.AuthController.Login)
	router.Post("/register", pr.AuthController.Register)
	router.Post("/whoami", pr.Security.Authentication, pr.AuthController.Whoami)
}
