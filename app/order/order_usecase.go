package order

import (
	"github.com/lathief/learn-fiber-go/app/dtos"
	"github.com/lathief/learn-fiber-go/app/models"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type orderUseCase struct {
	ProductRepo  repositories.ProductRepository
	OrderRepo    repositories.OrderRepository
	OrderProduct repositories.OrderProductRepository
}

type OrderUseCase interface {
	GetAllOrders() handlers.ReturnResponse
	GetOrderById(id int) handlers.ReturnResponse
	CreateOrder(order dtos.OrderReqDTO) handlers.ReturnResponse
	UpdateOrder(id int, product dtos.OrderDTO) handlers.ReturnResponse
	DeleteOrder(id int) handlers.ReturnResponse
}

func (ou orderUseCase) GetAllOrders() handlers.ReturnResponse {
	orders, err := ou.OrderRepo.GetAll()
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var ordersDTO []dtos.AllOrderDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, dtos.AllOrderDTO{
			ID:        order.ID,
			User:      order.UserId,
			Status:    order.Status,
			OrderDate: order.OrderDate,
		})
	}

	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    ordersDTO,
	}
}

func (ou orderUseCase) GetOrderById(id int) handlers.ReturnResponse {
	products, order, err := ou.OrderRepo.GetById(int64(id))
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	var sliceProducts []dtos.ProductDTO
	for _, item := range products {
		sliceProducts = append(sliceProducts, dtos.ProductDTO{
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	res := dtos.OrderDTO{
		ID:        order.ID,
		UserId:    order.UserId,
		Products:  sliceProducts,
		Status:    order.Status,
		OrderDate: order.OrderDate,
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
		Data:    res,
	}
}

func (ou orderUseCase) CreateOrder(order dtos.OrderReqDTO) handlers.ReturnResponse {
	var orderSave models.Order
	orderSave.Status = order.Status
	orderSave.UserId = order.UserId
	err := ou.OrderRepo.Create(orderSave, order.ProductsId)
	if err != nil {
		return handlers.ReturnResponse{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
	}
	return handlers.ReturnResponse{
		Code:    200,
		Message: "Success",
	}
}

func (ou orderUseCase) UpdateOrder(id int, product dtos.OrderDTO) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (ou orderUseCase) DeleteOrder(id int) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}
