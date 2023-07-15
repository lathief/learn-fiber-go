package cart

import (
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type cartUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
}
type CartUseCase interface {
	GetCartById(id int) (handlers.ReturnResponse, error)
	CreateCart(cart dtos.CartDTO) handlers.ReturnResponse
	UpdateCart(id int, category dtos.CartDTO) handlers.ReturnResponse
	Delete(id int) handlers.ReturnResponse
	GetAllCarts() handlers.ReturnResponse
}

func (cu *cartUseCase) GetCartById(id int) (handlers.ReturnResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (cu *cartUseCase) CreateCart(cart dtos.CartDTO) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (cu *cartUseCase) UpdateCart(id int, category dtos.CartDTO) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (cu *cartUseCase) Delete(id int) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (cu *cartUseCase) GetAllCarts() handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}
