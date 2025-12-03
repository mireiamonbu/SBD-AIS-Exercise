package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// todo
	// 1. List drinks
	// 2. Order a few drinks
	// 3. Order more drinks
	// 4. Get order total
	//
	// print responses after each call

	ctx := context.Background()

	// 1. List drinks
	fmt.Println("Requesting drinks...")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Avaliable drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d name:%q price:%d description:%q\n", d.Id, d.Name, d.Price, d.Description)
	}

	// 2. Order a few drinks
	fmt.Println("Ordering drinks...")
	Orders := []struct {
		id       int32
		quantity int32
		name     string
	}{
		{1, 2, "Spritzer"},
		{2, 2, "Beer"},
		{3, 2, "Coffee"},
	}

	for _, ord := range Orders {
		fmt.Printf("\t> Ordering: %d x %s\n", ord.quantity, ord.name)
		_, err := c.client.OrderDrink(ctx, &pb.OrderRequest{
			DrinkId:  ord.id,
			Quantity: ord.quantity,
		})
		if err != nil {
			return err
		}
	}

	// 3. Order more drinks
	fmt.Println("Ordering more drinks...")
	Orders2 := []struct {
		id       int32
		quantity int32
		name     string
	}{
		{1, 6, "Spritzer"},
		{2, 6, "Beer"},
		{3, 6, "Coffee"},
	}

	for _, ord := range Orders2 {
		fmt.Printf("\t> Ordering: %d x %s\n", ord.quantity, ord.name)
		_, err := c.client.OrderDrink(ctx, &pb.OrderRequest{
			DrinkId:  ord.id,
			Quantity: ord.quantity,
		})
		if err != nil {
			return err
		}
	}

	// 4. Get order total
	fmt.Println("Getting the bill:")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	total := map[int32]int32{}
	for _, ord := range ordersResp.Orders {
		total[ord.DrinkId] += ord.Quantity
	}

	name := map[int32]string{}
	for _, drink := range drinksResp.Drinks {
		name[drink.Id] = drink.Name
	}

	for id, quant := range total {
		fmt.Printf("\t> Total: %d x %s\n", quant, name[id])
	}

	return nil
}
