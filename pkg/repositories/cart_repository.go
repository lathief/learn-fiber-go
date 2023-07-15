package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type CartRepository interface {
	Create(ctx context.Context, cart models.Cart) error
	GetAll(ctx context.Context) ([]models.Cart, error)
	GetById(ctx context.Context, id int64) (models.Cart, error)
	GetByUserId(ctx context.Context, id int64) (models.Cart, error)
	AddProductsInCart(ctx context.Context, Items map[string]int64, userId int64) error
	DeleteAllProductsInCart(ctx context.Context, carId int64) error
}

func NewCartRepository(DB *sqlx.DB) CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

type cartRepository struct {
	DB *sqlx.DB
}

func (c *cartRepository) Create(ctx context.Context, cart models.Cart) error {
	res, err := c.DB.NamedExecContext(ctx, "INSERT INTO cart (user_id) VALUES (:user_id)", cart)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func (c *cartRepository) GetAll(ctx context.Context) ([]models.Cart, error) {
	var carts []models.Cart
	err := c.DB.SelectContext(ctx, &carts, `SELECT * FROM cart`)
	return carts, err
}

func (c *cartRepository) GetById(ctx context.Context, id int64) (models.Cart, error) {
	var cart models.Cart
	err := c.DB.GetContext(ctx, &cart, `SELECT * FROM cart WHERE id = ?`, id)
	return cart, err
}

func (c *cartRepository) GetByUserId(ctx context.Context, id int64) (models.Cart, error) {
	var cart models.Cart
	err := c.DB.GetContext(ctx, &cart, `SELECT * FROM cart WHERE user_id = ?`, id)
	return cart, err
}

func (c *cartRepository) AddProductsInCart(ctx context.Context, Items map[string]int64, cartId int64) error {
	// TODO: Memasukkan quantity dan productId ke dalam db berdasarkan cartID
	panic("implement me")
}
func (c *cartRepository) DeleteAllProductsInCart(ctx context.Context, carId int64) error {
	//TODO implement me
	panic("implement me")
}
