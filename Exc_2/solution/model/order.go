package model

type Order struct {
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
	DrinkID   uint64 `json:"drink_id"` // foreign key
	// todo Add fields: CreatedAt (time.Time), Amount with suitable types
	// todo json attributes need to be snakecase, i.e. name, created_at, my_variable, ..
}
