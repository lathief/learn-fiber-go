package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/app/models"
)

type orderProductRepository struct {
	DB *sqlx.DB
}
type OrderProductRepository interface {
	UpdateProductsInOrder(product models.Category) error
	DeleteProductInOrder(productId int64) error
}

func NewOrderProductRepository(DB *sqlx.DB) OrderProductRepository {
	return &orderProductRepository{
		DB: DB,
	}
}

func (op *orderProductRepository) UpdateProductsInOrder(product models.Category) error {
	//TODO implement me
	panic("implement me")
}
func (op *orderProductRepository) DeleteProductInOrder(productId int64) error {
	//TODO implement me
	panic("implement me")
}
