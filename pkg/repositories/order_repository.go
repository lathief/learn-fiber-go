package repositories

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/app/models"
)

type orderRepository struct {
	DB *sqlx.DB
}

type OrderRepository interface {
	Create(order models.Order, productsId []int64) error
	GetAll() ([]models.Order, error)
	GetById(id int64) (products []models.Product, order models.Order, err error)
	GetAllByProductId(productId int64) ([]models.Order, error)
	GetAllByCustomerId(customerId int64) ([]models.Order, error)
	Update(order models.Order) error
	Delete(id int64) error
}

func NewOrderRepository(DB *sqlx.DB) OrderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (o *orderRepository) Create(order models.Order, productsId []int64) error {
	tx, err := o.DB.Beginx()
	if err != nil {
		return err
	}
	res, err := tx.NamedExec("INSERT INTO `order` (user_id, status) VALUES (:user_id, :status)", order)
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
	for _, i := range productsId {
		res, err = tx.Exec("INSERT INTO order_product (order_id, product_id) VALUES (?, ?)", orderId, i)
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
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (o *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := o.DB.Select(&orders, "SELECT * FROM `order`")
	return orders, err
}
func (o *orderRepository) GetById(id int64) (products []models.Product, order models.Order, err error) {
	rows, err := o.DB.Query(
		"SELECT orders.*, products.* FROM order_product op INNER JOIN product AS products on op.product_id = products.id INNER JOIN `order` AS orders on op.order_id = orders.id WHERE orders.id = ?", id)
	if err != nil {
		return nil, models.Order{}, err
	}
	for rows.Next() {
		var tmpProduct models.Product
		err = rows.Scan(&order.ID, &order.UserId, &order.Status, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt, &tmpProduct.ID, &tmpProduct.Name, &tmpProduct.Price, &tmpProduct.Description, &tmpProduct.CategoryId, &tmpProduct.CreatedAt, &tmpProduct.UpdatedAt)
		if err != nil {
			return nil, models.Order{}, err
		}
		products = append(products, tmpProduct)
	}
	return products, order, err
}
func (o *orderRepository) GetAllByProductId(productId int64) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}
func (o *orderRepository) GetAllByCustomerId(customerId int64) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}
func (o *orderRepository) Update(order models.Order) error {
	//TODO implement me
	panic("implement me")
}
func (o *orderRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
