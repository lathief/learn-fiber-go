package repositories

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type productRepository struct {
	DB *sqlx.DB
}

type ProductRepository interface {
	Create(ctx context.Context, product models.Product) error
	GetAll(ctx context.Context) ([]models.Product, error)
	GetById(ctx context.Context, id int64) (models.Product, error)
	GetByIds(ctx context.Context, id []int64) ([]models.Product, error)
	GetAllByCategoryId(ctx context.Context, categoryId int64) ([]models.Product, error)
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, id int64) error
}

func NewProductRepository(DB *sqlx.DB) ProductRepository {
	return &productRepository{
		DB: DB,
	}
}
func (p *productRepository) Create(ctx context.Context, product models.Product) error {
	res, err := p.DB.NamedExecContext(ctx, "INSERT INTO product (name, price, description, category_id) VALUES (:name, :price, :description, :category_id)", product)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (p *productRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := p.DB.SelectContext(ctx, &products, `SELECT * FROM product`)
	return products, err
}

func (p *productRepository) GetById(ctx context.Context, id int64) (models.Product, error) {
	var product models.Product
	err := p.DB.GetContext(ctx, &product,
		`SELECT * FROM product WHERE product.id = ?`, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func (p *productRepository) Update(ctx context.Context, product models.Product) error {
	res, err := p.DB.NamedExecContext(ctx,
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

func (p *productRepository) Delete(ctx context.Context, id int64) error {
	res, err := p.DB.ExecContext(ctx, "DELETE FROM product WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (p *productRepository) GetAllByCategoryId(ctx context.Context, categoryId int64) ([]models.Product, error) {
	var product []models.Product
	err := p.DB.SelectContext(ctx, &product,
		`SELECT * FROM product WHERE product.category_id = ?`, categoryId)
	if err != nil {
		return nil, err
	}
	return product, err
}
func (p *productRepository) GetByIds(ctx context.Context, id []int64) ([]models.Product, error) {
	var products []models.Product
	var product models.Product
	query, args, err := sqlx.In("SELECT id, name, price, description FROM product WHERE id IN (?);", id)
	if err != nil {
		return nil, err
	}
	query = p.DB.Rebind(query)
	rows, err := p.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, err
}
