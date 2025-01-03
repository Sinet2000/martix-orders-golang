package order

import (
	"context"

	"github.com/Sinet2000/Martix-Orders-Go/proto"
)

type OrderServiceServer struct {
	proto.UnimplementedOrderServiceServer
	useCase OrderUseCase
}

func NewOrderServiceServer(useCase OrderUseCase) *OrderServiceServer {
	return &OrderServiceServer{useCase: useCase}
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	order := &Order{
		CustomerID:    req.CustomerId,
		Status:        "Pending",
		Total:         req.TotalPrice,
		PaymentMethod: req.PaymentMethod,
	}

	err := s.useCase.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return &proto.OrderResponse{
		Order: &proto.Order{
			Id:         int32(order.ID),
			CustomerId: order.CustomerID,
			Status:     order.Status,
			Total:      order.Total,
			CreatedAt:  order.CreatedAt,
		},
	}, nil
}

func (s *OrderServiceServer) GetOrderById(ctx context.Context, req *proto.GetOrderByIdRequest) (*proto.OrderResponse, error) {
	order, err := s.useCase.GetByCustomerID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.OrderResponse{
		Order: &proto.Order{
			Id:         int32(order.ID),
			CustomerId: order.CustomerID,
			Status:     order.Status,
			Total:      order.Total,
			CreatedAt:  order.CreatedAt,
		},
	}, nil
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	orders, err := s.useCase.GetAll(ctx, int(req.PageIndex), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	var protoOrders []*proto.Order
	for _, order := range orders {
		protoOrders = append(protoOrders, &proto.Order{
			Id:         int32(order.ID),
			CustomerId: order.CustomerID,
			Status:     order.Status,
			Total:      order.Total,
			CreatedAt:  order.CreatedAt,
		})
	}

	return &proto.ListOrdersResponse{
		Orders:     protoOrders,
		TotalCount: int32(len(orders)),
	}, nil
}
