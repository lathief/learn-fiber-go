package repositories

import (
	"fmt"
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
func (p productRepository) Create(product models.Product) error {
	res, err := p.DB.NamedExec("INSERT INTO product (name, price, description) VALUES (:name, :price, :description)", product)
	rowsAffected, _ := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}

func (p productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := p.DB.Select(&products, `SELECT * FROM product`)
	return products, err
}

func (p productRepository) GetById(id int64) (models.Product, error) {
	var product models.Product
	err := p.DB.Get(&product, `SELECT * FROM product WHERE id = ?`, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func (p productRepository) Update(product models.Product) error {
	fmt.Println(product.UpdatedAt)
	res, err := p.DB.NamedExec(
		"UPDATE product SET name=:name, price=:price, description=:description WHERE id=:id",
		product)
	rowsAffected, _ := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}

func (p productRepository) Delete(id int64) error {
	res, err := p.DB.NamedExec("DELETE FROM product WHERE id=:id", id)
	rowsAffected, _ := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}
