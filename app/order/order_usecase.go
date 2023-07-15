package order

import (
	"context"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type orderUseCase struct {
	ProductRepo repositories.ProductRepository
	OrderRepo   repositories.OrderRepository
}

type OrderUseCase interface {
	GetAllOrders(ctx context.Context) (ordersDTO []dtos.AllOrderDTO, err error)
	GetOrderById(ctx context.Context, id int) (orderDTO dtos.OrderDTO, err error)
	CreateOrder(ctx context.Context, order dtos.OrderReqDTO) (err error)
	UpdateOrder(ctx context.Context, id int, product dtos.OrderDTO) (err error)
	DeleteOrder(ctx context.Context, id int) (err error)
}

func (ou *orderUseCase) GetAllOrders(ctx context.Context) (ordersDTO []dtos.AllOrderDTO, err error) {
	orders, err := ou.OrderRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		ordersDTO = append(ordersDTO, dtos.AllOrderDTO{
			ID:         order.ID,
			User:       order.UserId,
			Status:     order.Status,
			TotalPrice: order.TotalPrice,
			OrderDate:  order.OrderDate,
		})
	}
	return ordersDTO, err
}

func (ou *orderUseCase) GetOrderById(ctx context.Context, id int) (orderDTO dtos.OrderDTO, err error) {
	products, order, err := ou.OrderRepo.GetById(ctx, int64(id))
	if err != nil {
		return dtos.OrderDTO{}, err
	}
	var sliceProducts []dtos.ProductDTO
	for _, item := range products {
		sliceProducts = append(sliceProducts, dtos.ProductDTO{
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	orderDTO = dtos.OrderDTO{
		ID:         order.ID,
		UserId:     order.UserId,
		Products:   sliceProducts,
		Status:     order.Status,
		TotalPrice: order.TotalPrice,
		OrderDate:  order.OrderDate,
	}
	return orderDTO, nil
}

func (ou *orderUseCase) CreateOrder(ctx context.Context, order dtos.OrderReqDTO) (err error) {
	var orderSave models.Order
	orderSave.Status = order.Status
	orderSave.UserId = order.UserId
	err = ou.OrderRepo.Create(ctx, orderSave, order.ProductsId)
	return err
}

func (ou *orderUseCase) UpdateOrder(ctx context.Context, id int, product dtos.OrderDTO) (err error) {
	//TODO implement me
	panic("implement me")
}

func (ou *orderUseCase) DeleteOrder(ctx context.Context, id int) (err error) {
	//TODO implement me
	panic("implement me")
}
