package db

import (
	"context"
	"ecommerce/util"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateInventory(t *testing.T) {
	var price pgtype.Numeric
	_ = price.Scan("50.99")
	var updatedAt pgtype.Timestamptz
	_ = updatedAt.Scan(time.Now())

	// Use a unique product name to avoid conflicts
	productName := fmt.Sprintf("Test Product %d", time.Now().UnixNano())

	args := CreateInventoryParams{
		ProductName: productName,
		Price:       price,
		Currency:    "USD",
		Stock:       50,
		UpdatedAt:   updatedAt,
	}

	inv, err := testQueries.CreateInventory(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, inv)
	require.Equal(t, args.Price, inv.Price)
	require.Equal(t, args.Currency, inv.Currency)
	require.Equal(t, args.Stock, inv.Stock)
	require.Equal(t, args.ProductName, inv.ProductName)

	require.NotZero(t, inv.ProductID)
	require.NotZero(t, inv.UpdatedAt)

	// Clean up: delete the created inventory item
	err = testQueries.DeleteInventory(context.Background(), inv.ProductID)
	require.NoError(t, err)
}

func CreateInventory(t *testing.T) Inventory {

	/*args := CreateInventoryParams{
		Price:     price,
		Currency:  "USD",
		Stock:     50,
		UpdatedAt: updatedAt,
	}*/

	/*inv, err := testQueries.CreateInventory(context.Background(), CreateInventoryParams{
		ProductName: util.RandomProduct(),
		Price:       price,
		Currency:    util.RandomCurrency(),
		Stock:       int32(util.RandomInt(0, 200)),
		UpdatedAt:   pgtype.Timestamptz{},
	})*/
	//require.NoError(t, err)
	//require.NotEmpty(t, inv)
	/*require.Equal(t, args.Price, inv.Price)
	require.Equal(t, args.Currency, inv.Currency)
	require.Equal(t, args.Stock, inv.Stock)

	require.NotZero(t, inv.ProductID)
	require.NotZero(t, inv.UpdatedAt)*/

	// Clean up: delete the created inventory item
	//err = testQueries.DeleteInventory(context.Background(), inv.ProductID)
	//require.NoError(t, err)
	// First create a product in inventory
	var price pgtype.Numeric
	_ = price.Scan("999.99")

	inv, err := testQueries.CreateInventory(context.Background(), CreateInventoryParams{
		ProductName: "iPhone 14 Pro",
		Price:       price,
		Currency:    "USD",
		Stock:       10,
		UpdatedAt:   pgtype.Timestamptz{},
	})
	require.NoError(t, err)
	require.NotEmpty(t, inv)

	t.Logf("Inserted product: %s for %v %s", util.RandomProduct(), inv.Price, inv.Currency)
	return inv

}
func TestFakeCreateInventory(t *testing.T) {
	price := pgtype.Numeric{
		Int:   big.NewInt(12345), // Represents 123.45 as an integer (unscaled)
		Exp:   -2,                // -2 means two decimal places
		Valid: true,              // Marks the value as present (not null)
	}

	var createdProducts []int32 // Track created product IDs for cleanup

	for i := 0; i <= 10; i++ {
		// Use index and timestamp to ensure uniqueness
		productName := fmt.Sprintf("%s_Test_%d_%d", util.RandomProduct(), i, time.Now().UnixNano())

		inv, err := testQueries.CreateInventory(context.Background(), CreateInventoryParams{
			ProductName: productName,
			Price:       price,
			Currency:    util.RandomCurrency(),
			Stock:       int32(util.RandomInt(0, 1000)),
			UpdatedAt:   pgtype.Timestamptz{},
		})
		require.NoError(t, err)
		require.NotEmpty(t, inv)
		createdProducts = append(createdProducts, inv.ProductID)
		t.Logf("Inserted product: %s for %v %s with stock: %d", inv.ProductName, inv.Price, inv.Currency, inv.Stock)
	}

	// Clean up all created products
	for _, productID := range createdProducts {
		err := testQueries.DeleteInventory(context.Background(), productID)
		require.NoError(t, err)
	}
}

func TestDeleteInventory(t *testing.T) {
	//account1 := CreateInventory()
	err := testQueries.DeleteInventory(context.Background(), 0)
	require.NoError(t, err)
	// Clean up
	//err = testQueries.DeleteInventory(context.Background(), inv.ProductID)
	//require.NoError(t, err)
}
