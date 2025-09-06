package orders

import (
	"context"
	db "ecommerce/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService struct {
	*db.Queries
	db *pgxpool.Pool //this is the database connection
}

func NewOrderService(dbConn *pgxpool.Pool) *OrderService {
	//run the connection
	return &OrderService{
		Queries: db.New(dbConn),
		db:      dbConn,
	}
}

func (o *OrderService) PlaceOrder(ctx context.Context, ProductName string, Quantity int) (db.Order, error) {
	// the database logic:
	//let us start the connection

	//find product by name and specify the quantity

	//create a pending status
	//update inventory so that the stock - quantity
	//later update it to shipping
	return db.Order{}, nil
}
