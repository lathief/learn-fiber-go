package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type productRepository struct {
	DB *sqlx.DB
}

type ProductRepository interface {
	Create(product models.Product) error
	GetAll() ([]models.Product, error)
	GetById(id int64) (models.Product, error)
	GetAllByCategoryId(categoryId int64) ([]models.Product, error)
	GetAllProductInOrderById(id int64) ([]models.Product, models.Order, error)
	Update(product models.Product) error
	Delete(id int64) error
}

func NewProductRepository(DB *sqlx.DB) ProductRepository {
	return &productRepository{
		DB: DB,
	}
}
func (p *productRepository) Create(product models.Product) error {
	res, err := p.DB.NamedExec("INSERT INTO product (name, price, description, category_id) VALUES (:name, :price, :description, :category_id)", product)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := p.DB.Select(&products, `SELECT * FROM product`)
	return products, err
}

func (p *productRepository) GetById(id int64) (models.Product, error) {
	var product models.Product
	err := p.DB.Get(&product,
		`SELECT * FROM product WHERE product.id = ?`, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func (p *productRepository) Update(product models.Product) error {
	res, err := p.DB.NamedExec(
		`UPDATE product SET name=:name, price=:price, 
				description= CASE WHEN :description IS NOT NULL AND LENGTH(:description) > 0 THEN :description ELSE description END,
				category_id=:category_id WHERE id = :id`,
		product)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (p *productRepository) Delete(id int64) error {
	res, err := p.DB.Exec("DELETE FROM product WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (p *productRepository) GetAllByCategoryId(categoryId int64) ([]models.Product, error) {
	var product []models.Product
	err := p.DB.Select(&product,
		`SELECT * FROM product WHERE product.category_id = ?`, categoryId)
	if err != nil {
		return nil, err
	}
	return product, err
}
func (p *productRepository) GetAllProductInOrderById(id int64) ([]models.Product, models.Order, error) {
	var products []models.Product
	var order models.Order
	rows, err := p.DB.Query(
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
