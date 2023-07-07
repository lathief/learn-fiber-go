package order

import (
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/repositories"
)

type orderUseCase struct {
	ProductRepo  repositories.ProductRepository
	CategoryRepo repositories.CategoryRepository
	OrderRepo    repositories.OrderRepository
}

type OrderUseCase interface {
	GetAllOrders() handlers.ReturnResponse
	GetOrderById(id int) handlers.ReturnResponse
	CreateOrder(order dtos.OrderDTO) handlers.ReturnResponse
	UpdateOrder(id int, product dtos.OrderDTO) handlers.ReturnResponse
	DeleteOrder(id int) handlers.ReturnResponse
}

func (ou orderUseCase) GetAllOrders() handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (ou orderUseCase) GetOrderById(id int) handlers.ReturnResponse {
	products, order, err := ou.ProductRepo.GetAllProductInOrderById(int64(id))
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

func (ou orderUseCase) CreateOrder(order dtos.OrderDTO) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (ou orderUseCase) UpdateOrder(id int, product dtos.OrderDTO) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}

func (ou orderUseCase) DeleteOrder(id int) handlers.ReturnResponse {
	//TODO implement me
	panic("implement me")
}
