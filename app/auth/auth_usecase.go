package auth

import (
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type authUseCase struct {
	UserRepo repositories.UserRepository
}

type AuthUseCase interface {
	Register(newUser dtos.LoginDTO) error
	Login(auth dtos.LoginDTO) (dtos.TokenAuth, error)
	Whoami(token string) (dtos.UserDTO, error)
}

func (a *authUseCase) Register(newUser dtos.LoginDTO) error {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) Login(auth dtos.LoginDTO) (dtos.TokenAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) Whoami(token string) (dtos.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}
