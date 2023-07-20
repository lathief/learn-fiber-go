package auth

import (
	"context"
	"errors"
	"github.com/lathief/learn-fiber-go/pkg/constant"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
	"github.com/lathief/learn-fiber-go/pkg/utils"
	"strconv"
)

type authUseCase struct {
	UserRepo repositories.UserRepository
	RoleRepo repositories.RoleRepository
	CartRepo repositories.CartRepository
}

type AuthUseCase interface {
	Register(ctx context.Context, newUser dtos.RegisterDTO) (err error)
	Login(ctx context.Context, auth dtos.LoginDTO) (token dtos.TokenAuth, err error)
	Whoami(ctx context.Context, userId string) (dtos.UserDTO, error)
}

func (a *authUseCase) Register(ctx context.Context, newUser dtos.RegisterDTO) (err error) {
	var errMsg string
	var role models.Role
	existUser, _ := a.UserRepo.GetByUsername(ctx, newUser.Username)
	if existUser.Username != "" {
		errMsg = utils.JoinStringWithComma(errMsg, constant.ErrUsernameExists.Error())
	}
	existUser, _ = a.UserRepo.GetByEmail(ctx, newUser.Email)
	if existUser.Email != "" {
		errMsg = utils.JoinStringWithComma(errMsg, constant.ErrEmailExists.Error())
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	if newUser.RoleName == "" {
		role, err = a.RoleRepo.GetByName(ctx, constant.ROLE_CUSTOMER)
		if err != nil {
			return constant.ErrRoleNotFound
		}
	} else {
		role, err = a.RoleRepo.GetByName(ctx, newUser.RoleName)
		if err != nil {
			return constant.ErrRoleNotFound
		}
	}
	userSave := models.User{
		Username:  newUser.Username,
		Password:  utils.HashPass(newUser.Password),
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		RoleId:    role.ID,
	}

	id, err := a.UserRepo.Create(ctx, userSave)
	if err != nil {
		return err
	}
	err = a.CartRepo.Create(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *authUseCase) Login(ctx context.Context, auth dtos.LoginDTO) (token dtos.TokenAuth, err error) {
	var getUser models.User
	if auth.Email == "" {
		getUser, err = a.UserRepo.GetByUsername(ctx, auth.Username)
		if err != nil {
			return dtos.TokenAuth{}, constant.ErrUserNotFound
		}
	}
	if auth.Username == "" {
		getUser, err = a.UserRepo.GetByEmail(ctx, auth.Email)
		if err != nil {
			return dtos.TokenAuth{}, constant.ErrUserNotFound
		}
	}
	match := utils.ComparePass([]byte(getUser.Password), []byte(auth.Password))
	if match == false {
		return dtos.TokenAuth{}, constant.ErrPasswordNotMatch
	}
	role, err := a.RoleRepo.GetById(ctx, getUser.RoleId)
	if err != nil {
		return dtos.TokenAuth{}, constant.ErrRoleNotFound
	}
	accessToken := utils.GenerateToken(int(getUser.ID), getUser.Username, role.RoleName)
	refreshToken, err := utils.RefreshToken(accessToken)
	if err != nil {
		return dtos.TokenAuth{}, err
	}
	return dtos.TokenAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUseCase) Whoami(ctx context.Context, userId string) (dtos.UserDTO, error) {
	var getUser models.User
	id, err := strconv.Atoi(userId)
	if err != nil {
		return dtos.UserDTO{}, err
	}
	getUser, err = a.UserRepo.GetById(ctx, int64(id))
	if err != nil {
		return dtos.UserDTO{}, err
	}
	return dtos.UserDTO{
		ID:        getUser.ID,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Email:     getUser.Email,
	}, nil
}
