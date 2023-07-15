package repositories

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type CategoryRepository interface {
	Create(ctx context.Context, category models.Category) error
	GetAll(ctx context.Context) ([]models.Category, error)
	GetById(ctx context.Context, id int64) (models.Category, error)
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id int64) error
}

func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		DB: DB,
	}
}

type categoryRepository struct {
	DB *sqlx.DB
}

func (c *categoryRepository) Create(ctx context.Context, category models.Category) error {
	res, err := c.DB.NamedExecContext(ctx, "INSERT INTO category (name, description) VALUES (:name, :description)", category)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func (c *categoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var category []models.Category
	err := c.DB.SelectContext(ctx, &category, `SELECT * FROM category`)
	return category, err
}

func (c *categoryRepository) GetById(ctx context.Context, id int64) (models.Category, error) {
	var category models.Category
	err := c.DB.GetContext(ctx, &category, `SELECT * FROM category WHERE id = ?`, id)
	return category, err
}

func (c *categoryRepository) Update(ctx context.Context, category models.Category) error {
	res, err := c.DB.NamedExecContext(ctx,
		`UPDATE category SET name=:name, 
				description= CASE WHEN :description IS NOT NULL AND LENGTH(:description) > 0 THEN :description ELSE description END
				WHERE id = :id`,
		category)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (c *categoryRepository) Delete(ctx context.Context, id int64) error {
	res, err := c.DB.ExecContext(ctx, "DELETE FROM category WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
