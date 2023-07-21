package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type CartRepository interface {
	Create(ctx context.Context, userId int64) error
	GetByUserId(ctx context.Context, id int64) (models.Cart, error)
	GetItemsInCart(ctx context.Context, cartId int64) ([]models.CartItems, error)
	AddProductsInCart(ctx context.Context, cartItem models.CartItems) error
	DeleteProductsInCart(ctx context.Context, cartId int64, productId int64) error
}

func NewCartRepository(DB *sqlx.DB) CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

type cartRepository struct {
	DB *sqlx.DB
}

func (c *cartRepository) Create(ctx context.Context, userId int64) error {
	res, err := c.DB.ExecContext(ctx, "INSERT INTO cart (user_id) VALUES (?)", userId)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func (c *cartRepository) GetById(ctx context.Context, id int64) (models.Cart, error) {
	var cart models.Cart
	err := c.DB.GetContext(ctx, &cart, `SELECT * FROM cart WHERE id = ?`, id)
	return cart, err
}
func (c *cartRepository) GetItemsInCart(ctx context.Context, cartId int64) ([]models.CartItems, error) {
	var cart []models.CartItems
	err := c.DB.SelectContext(ctx, &cart, `SELECT * FROM cart_items WHERE cart_id = ?`, cartId)
	return cart, err
}
func (c *cartRepository) GetByUserId(ctx context.Context, id int64) (models.Cart, error) {
	var cart models.Cart
	err := c.DB.GetContext(ctx, &cart, `SELECT * FROM cart WHERE user_id = ?`, id)
	return cart, err
}

func (c *cartRepository) AddProductsInCart(ctx context.Context, cartItem models.CartItems) error {
	res, err := c.DB.NamedExecContext(ctx, "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (:cart_id, :product_id, :quantity)", cartItem)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}
func (c *cartRepository) DeleteProductsInCart(ctx context.Context, cartId int64, productId int64) error {
	res, err := c.DB.ExecContext(ctx, "DELETE FROM cart_items WHERE cart_id = ? AND product_id = ?", cartId, productId)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}
