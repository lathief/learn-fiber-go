package repositories

import (
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
	rowsAffected, _ := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}

func (c *categoryRepository) GetAll() ([]models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) GetById(id int64) (models.Category, error) {
	var category models.Category
	err := c.DB.Get(&category, `SELECT category.* FROM product AS products INNER JOIN category ON category.id = products.category_id WHERE category_id = ?`, id)
	return category, err
}

func (c *categoryRepository) Update(category models.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
