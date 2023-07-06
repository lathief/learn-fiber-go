package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/models"
)

type CategoryRepository interface {
	Create(category models.Category) error
	GetAll() ([]models.Category, error)
	GetById(id int64) (models.Category, error)
	Update(product models.Category) error
	Delete(id int64) error
}

func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		DB: DB,
	}
}

type categoryRepository struct {
	DB *sqlx.DB
}

func (c *categoryRepository) Create(category models.Category) error {
	res, err := c.DB.NamedExec("INSERT INTO category (name, description) VALUES (:name, :description)", category)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}

func (c *categoryRepository) GetAll() ([]models.Category, error) {
	var category []models.Category
	err := c.DB.Select(&category, `SELECT * FROM category`)
	return category, err
}

func (c *categoryRepository) GetById(id int64) (models.Category, error) {
	var category models.Category
	err := c.DB.Get(&category, `SELECT * FROM category WHERE id = ?`, id)
	return category, err
}

func (c *categoryRepository) Update(category models.Category) error {
	res, err := c.DB.NamedExec(
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

func (c *categoryRepository) Delete(id int64) error {
	res, err := c.DB.Exec("DELETE FROM category WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
