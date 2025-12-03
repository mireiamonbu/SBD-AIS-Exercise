package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks map[int32]*pb.Drink
	orders []*pb.Order
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions
func (s *GRPCService) Init() {
	s.drinks = map[int32]*pb.Drink{
		1: {Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
		2: {Id: 2, Name: "Beer", Price: 5, Description: "Hagenberg Gold"},
		3: {Id: 3, Name: "Coffee", Price: 3, Description: "Mifare isn't that secure"},
	}

	s.orders = []*pb.Order{}

}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*wrapperspb.BoolValue, error) {

	// Save the order
	s.orders = append(s.orders, &pb.Order{
		DrinkId:  req.DrinkId,
		Quantity: req.Quantity,
	})

	return &wrapperspb.BoolValue{Value: true}, nil
}

func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.GetDrinkResponse, error) {
	drinks := make([]*pb.Drink, 0, len(s.drinks))

	for _, d := range s.drinks {
		drinks = append(drinks, d)
	}

	return &pb.GetDrinkResponse{Drinks: drinks}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	return &pb.GetOrdersResponse{Orders: s.orders}, nil
}
