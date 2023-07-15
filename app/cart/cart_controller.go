package cart

import "github.com/gofiber/fiber/v2"

type cartController struct {
	CartUseCase CartUseCase
}
type CartController interface {
	GetCartById(ctx *fiber.Ctx) error
	CreateCart(ctx *fiber.Ctx) error
	GetAllCarts(ctx *fiber.Ctx) error
	UpdateCart(ctx *fiber.Ctx) error
	DeleteCart(ctx *fiber.Ctx) error
}

func (cc *cartController) GetCartById(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (cc *cartController) CreateCart(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (cc *cartController) GetAllCarts(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (cc *cartController) UpdateCart(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (cc *cartController) DeleteCart(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
