package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type orderRepository struct {
	DB *sqlx.DB
}
type OrderRepository interface {
	Create(order models.Order) error
	GetAll() ([]models.Order, error)
	GetById(id int64) (models.Order, error)
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
func (o *orderRepository) Create(order models.Order) error {
	res, err := o.DB.NamedExec("INSERT INTO `order` (customer_id, product_id) VALUES (:customer_id, :product_id)", order)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (o *orderRepository) GetAll() ([]models.Order, error) {
	var category []models.Order
	err := o.DB.Select(&category, `SELECT * FROM "order"`)
	return category, err
}
func (o *orderRepository) GetById(id int64) (models.Order, error) {
	var order models.Order
	err := o.DB.Get(&order, "SELECT * FROM `order` WHERE id = ?", id)
	return order, err
}
func (o *orderRepository) GetAllByProductId(productId int64) ([]models.Order, error) {
	var order []models.Order
	err := o.DB.Select(&order,
		`SELECT * FROM "order" WHERE product_id = ?`, productId)
	if err != nil {
		return nil, err
	}
	return order, err
}
func (o *orderRepository) GetAllByCustomerId(customerId int64) ([]models.Order, error) {
	var order []models.Order
	err := o.DB.Select(&order,
		`SELECT * FROM "order" WHERE customer_id = ?`, customerId)
	if err != nil {
		return nil, err
	}
	return order, err
}
func (o *orderRepository) Update(order models.Order) error {
	res, err := o.DB.NamedExec(
		`UPDATE "order" SET 
				status=CASE WHEN:status IS NOT NULL AND LENGTH(:status) > 0 THEN :status ELSE status END,
				product_id=:product_id WHERE id = :id`,
		order)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (o *orderRepository) Delete(id int64) error {
	res, err := o.DB.Exec(`DELETE FROM "order" WHERE id=?`, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
