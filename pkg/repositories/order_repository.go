package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type orderRepository struct {
	DB *sqlx.DB
}

type OrderRepository interface {
	Create(ctx context.Context, order models.Order, productItems []models.OrderItems) error
	GetAll(ctx context.Context) ([]models.Order, error)
	GetById(ctx context.Context, id int64) (products []models.Product, order models.Order, err error)
	GetAllByUserId(ctx context.Context, id int64) (models.Order, error)
}

func NewOrderRepository(DB *sqlx.DB) OrderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (o *orderRepository) Create(ctx context.Context, order models.Order, productItems []models.OrderItems) error {
	tx, err := o.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	res, err := tx.ExecContext(ctx, "INSERT INTO `order` (user_id, status) VALUES (?, ?)", order.UserId, order.Status)
	if err != nil {
		fmt.Println("Error inserting order")
		fmt.Println(err.Error())
		RollbackErr := tx.Rollback()
		if RollbackErr != nil {
			return RollbackErr
		}
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		RollbackErr := tx.Rollback()
		if RollbackErr != nil {
			return RollbackErr
		}
		return errors.New("insert order failed")
	}
	orderId, _ := res.LastInsertId()
	fmt.Println(orderId)
	for _, i := range productItems {
		res, err = tx.ExecContext(ctx, "INSERT INTO order_items (order_id, product_id, quantity) VALUES (?, ?, ?)", orderId, i.ProductId, i.Quantity)
		if err != nil {
			fmt.Println("Error inserting order product")
			RollbackErr := tx.Rollback()
			if RollbackErr != nil {
				return RollbackErr
			}
			return err
		}
		fmt.Println("Lanjut")
		rowsAffected, _ = res.RowsAffected()
		if rowsAffected == 0 {
			RollbackErr := tx.Rollback()
			if RollbackErr != nil {
				return RollbackErr
			}
			return errors.New("insert order product failed")
		}
	}
	fmt.Println("Lanjut")
	res, err = tx.ExecContext(ctx, `CALL COUNT_PRICE(?)`, orderId)
	if err != nil {
		fmt.Println("Error exec sp count items")
		RollbackErr := tx.Rollback()
		if RollbackErr != nil {
			return RollbackErr
		}
		return err
	}
	rowsAffected, _ = res.RowsAffected()
	if rowsAffected == 0 {
		RollbackErr := tx.Rollback()
		if RollbackErr != nil {
			return RollbackErr
		}
		return errors.New("Error exec sp count items")
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (o *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := o.DB.SelectContext(ctx, &orders, "SELECT * FROM `order`")
	return orders, err
}
func (o *orderRepository) GetById(ctx context.Context, id int64) (products []models.Product, order models.Order, err error) {
	rows, err := o.DB.QueryContext(ctx,
		"SELECT orders.*, products.* FROM order_items op INNER JOIN product AS products on op.product_id = products.id INNER JOIN `order` AS orders on op.order_id = orders.id WHERE orders.id = ?", id)
	if err != nil {
		return nil, models.Order{}, err
	}
	if !rows.Next() {
		return nil, models.Order{}, sql.ErrNoRows
	}
	for rows.Next() {
		var tmpProduct models.Product
		err = rows.Scan(&order.ID, &order.UserId, &order.Status, &order.TotalPrice, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt, &tmpProduct.ID, &tmpProduct.Name, &tmpProduct.Price, &tmpProduct.Description, &tmpProduct.CategoryId, &tmpProduct.CreatedAt, &tmpProduct.UpdatedAt)
		if err != nil {
			return nil, models.Order{}, err
		}
		products = append(products, tmpProduct)
	}
	return products, order, err
}
func (o *orderRepository) GetAllByUserId(ctx context.Context, id int64) (models.Order, error) {
	var order models.Order
	err := o.DB.GetContext(ctx, &order, "SELECT * FROM `order` WHERE id = ?", id)
	return order, err
}
