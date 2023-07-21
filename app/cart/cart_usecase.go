package cart

import (
	"context"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type cartUseCase struct {
	CartRepo    repositories.CartRepository
	ProductRepo repositories.ProductRepository
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
	var ids []int64
	for _, product := range cartProducts {
		ids = append(ids, product.ProductId)
	}
	products, err := cu.ProductRepo.GetByIds(ctx, ids)
	var items []dtos.Item
	for _, item := range cartProducts {
		items = append(items, dtos.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	for _, product := range products {
		for i, item := range items {
			if item.ProductId == product.ID {
				items[i].Name = product.Name
				items[i].Price = product.Price
				items[i].Description = product.Description
			}
		}
	}
	cartUser := dtos.CartDTO{
		ID:     getCart.ID,
		UserId: getCart.UserID,
		Items:  items,
	}
	return cartUser, nil
}
func (cu *cartUseCase) UpdateProductInCart(ctx context.Context, userId int, product dtos.Item) (err error) {
	getCart, err := cu.CartRepo.GetByUserId(ctx, int64(userId))
	if err != nil {
		return err
	}
	var cartItems = models.CartItems{
		CartId:    getCart.ID,
		ProductId: product.ProductId,
		Quantity:  product.Quantity,
	}
	err = cu.CartRepo.AddProductsInCart(ctx, cartItems)
	if err != nil {
		return err
	}
	return nil
}
func (cu *cartUseCase) DeleteProductsInCart(ctx context.Context, userId int, productid dtos.CartProductIdDTO) (err error) {
	getCart, err := cu.CartRepo.GetByUserId(ctx, int64(userId))
	if err != nil {
		return err
	}
	err = cu.CartRepo.DeleteProductsInCart(ctx, getCart.ID, int64(productid.ProductId))
	if err != nil {
		return err
	}
	return nil
}
