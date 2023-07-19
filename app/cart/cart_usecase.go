package cart

import (
	"context"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type cartUseCase struct {
	CartRepo repositories.CartRepository
}
type CartUseCase interface {
	GetCartByUserId(ctx context.Context, id int) (cart dtos.CartDTO, err error)
	UpdateProductInCart(ctx context.Context, userId int, product dtos.Item) (err error)
	CreateCart(ctx context.Context, userId int) (err error)
	DeleteProductsInCart(ctx context.Context, userId int, productid dtos.CartProductIdDTO) (err error)
}

func (cu *cartUseCase) CreateCart(ctx context.Context, userId int) (err error) {
	err = cu.CartRepo.Create(ctx, int64(userId))
	return err
}
func (cu *cartUseCase) GetCartByUserId(ctx context.Context, id int) (cart dtos.CartDTO, err error) {
	getCart, err := cu.CartRepo.GetByUserId(ctx, int64(id))
	if err != nil {
		return dtos.CartDTO{}, err
	}
	cartProducts, err := cu.CartRepo.GetItemsInCart(ctx, getCart.ID)
	if err != nil {
		return dtos.CartDTO{}, err
	}
	var items []dtos.Item
	for _, item := range cartProducts {
		items = append(items, dtos.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	cartUser := dtos.CartDTO{
		ID:     getCart.ID,
		UserId: getCart.UserID,
		Items:  items,
	}
	return cartUser, nil
}

func (cu *cartUseCase) DeleteProductsInCart(ctx context.Context, userId int, productid dtos.CartProductIdDTO) (err error) {
	//TODO implement me
	panic("implement me")
}
func (cu *cartUseCase) UpdateProductInCart(ctx context.Context, userId int, product dtos.Item) (err error) {
	//TODO implement me
	panic("implement me")
}
