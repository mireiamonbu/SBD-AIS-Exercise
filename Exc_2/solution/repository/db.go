package repository

import (
	"ordersystem/model"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	drinks := []model.Drink{
		{Description: "Normal water", ID: 1, Name: "Water", Price: 0.59},
		{Description: "Natural orange juice", ID: 2, Name: "Orange Juice", Price: 1.00},
		{Description: "Vermut", ID: 3, Name: "Martini", Price: 4.30},
		{Description: "Wine of selection", ID: 4, Name: "White Wine", Price: 2.50},
		{Description: "Wine of selection", ID: 5, Name: "Red Wine", Price: 3.00},
		{Description: "Normal beer", ID: 6, Name: "Beer", Price: 3.00},
	}
	// Init the drinks slice with some test data
	// drinks := ...

	// Init orders slice with some test data
	orders := []model.Order{
		{Amount: 3, CreatedAt: "16:00", DrinkID: 2},
		{Amount: 4, CreatedAt: "16:31", DrinkID: 1},
		{Amount: 2, CreatedAt: "17:00", DrinkID: 2},
		{Amount: 1, CreatedAt: "17:03", DrinkID: 5},
		{Amount: 3, CreatedAt: "17:28", DrinkID: 4},
		{Amount: 6, CreatedAt: "17:44", DrinkID: 5},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64

	totalledOrders := make(map[uint64]uint64)

	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += uint64(order.Amount)
	}

	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice

	db.orders = append(db.orders, *order)
}
