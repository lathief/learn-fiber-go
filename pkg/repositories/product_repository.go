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
	Update(product models.Product) error
	Delete(id int64) error
}

func NewProductRepository(DB *sqlx.DB) ProductRepository {
	return &productRepository{
		DB: DB,
	}
}
func (p *productRepository) Create(product models.Product) error {
	res, err := p.DB.NamedExec("INSERT INTO product (name, price, description, category_id) VALUES (:product_name, :product_price, :product_description, :product_category_id)", product)
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
	err := p.DB.Select(&products, `SELECT product.id AS product_id, product.name AS product_name, product.description AS product_description,
       			product.price AS product_price, product.category_id AS product_category_id, product.created_at AS product_created_at,
       			product.updated_at AS product_updated_at FROM product`)
	return products, err
}

func (p *productRepository) GetById(id int64) (models.Product, error) {
	var product models.Product
	err := p.DB.Get(&product,
		`SELECT product.id AS product_id, product.name AS product_name, product.description AS product_description,
       			product.price AS product_price, product.category_id AS product_category_id, product.created_at AS product_created_at,
       			product.updated_at AS product_updated_at, category.id AS 'product_category.category_id',
       			category.name AS 'product_category.category_name', category.description AS 'product_category.category_description',
       			category.created_at AS 'product_category.category_created_at', category.updated_at AS 'product_category.category_updated_at'
				FROM product
				INNER JOIN category ON category.id = product.category_id
				WHERE product.id = ?`, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func (p *productRepository) Update(product models.Product) error {
	res, err := p.DB.NamedExec(
		`UPDATE product SET name=:product_name, price=:product_price, 
				description= CASE WHEN :product_description IS NOT NULL AND LENGTH(:product_description) > 0 THEN :product_description ELSE description END,
				category_id=:product_category_id WHERE id = :product_id`,
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
