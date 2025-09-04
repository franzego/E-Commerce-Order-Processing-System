package db

import (
	"context"
	"ecommerce/util"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateOrderByProductName(t *testing.T) {
	// First create a product in inventory
	var price pgtype.Numeric
	_ = price.Scan("999.99")

	// Use a unique product name to avoid conflicts

	productName := util.RandomName()

	inv, err := testQueries.CreateInventory(context.Background(), CreateInventoryParams{
		ProductName: productName,
		Price:       price,
		Currency:    "USD",
		Stock:       10,
		UpdatedAt:   pgtype.Timestamptz{},
	})
	require.NoError(t, err)
	require.NotEmpty(t, inv)

	// Now create an order for this product
	args := CreateOrderByProducNameParams{
		ProductName: productName,
		Column2:     2,
	}

	ord, err := testQueries.CreateOrderByProducName(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, ord)
	require.Equal(t, inv.ProductID, ord.ProductID)
	require.Equal(t, int32(2), ord.Quantity)
	require.Equal(t, "pending", ord.Status)

	// Clean up: Delete the order first (due to foreign key constraint)
	err = testQueries.DeleteOrder(context.Background(), ord.ID)
	require.NoError(t, err)

	// Then delete the inventory
	err = testQueries.DeleteInventory(context.Background(), inv.ProductID)
	require.NoError(t, err)
}
