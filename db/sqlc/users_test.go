package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	timestamp := time.Now().UnixNano()
	arg := CreateUserParams{
		Username:     fmt.Sprintf("John@123_TestCreateUser_%d", timestamp),
		Email:        fmt.Sprintf("Freak@franz.com_TestCreateUser_%d", timestamp),
		PasswordHash: "12345",
	}

	use, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, use)
	require.NotEmpty(t, use.PasswordHash)
	require.NotZero(t, use.ID)

	err = testQueries.DeleteUser(context.Background(), use.ID)
	require.NoError(t, err)

}

func TestGetUserByEmail(t *testing.T) {
	// Create the user first before we try getting the user by email
	timestamp := time.Now().UnixNano()
	arg := CreateUserParams{
		Username:     fmt.Sprintf("John@123_TestGetUserByEmail_%d", timestamp),
		Email:        fmt.Sprintf("Freak@franz.com_TestGetUserByEmail_%d", timestamp),
		PasswordHash: "12345",
	}

	use, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, use)
	require.NotEmpty(t, use.PasswordHash)
	require.NotZero(t, use.ID)

	// we can now get the user by email
	email := fmt.Sprintf("Freak@franz.com_TestGetUserByEmail_%d", timestamp)
	use, err = testQueries.GetUserByEmail(context.Background(), email)
	require.NoError(t, err)
	require.NotEmpty(t, use)

	// Clean up: Delete the user
	err = testQueries.DeleteUser(context.Background(), use.ID)
	require.NoError(t, err)
}

func TestGetUserByUsername(t *testing.T) {
	// Create the user first before we try getting the user by username
	timestamp := time.Now().UnixNano()
	arg := CreateUserParams{
		Username:     fmt.Sprintf("John@123_TestGetUserByUsername_%d", timestamp),
		Email:        fmt.Sprintf("Freak@franz.com_TestGetUserByUsername_%d", timestamp),
		PasswordHash: "12345",
	}

	use, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, use)
	require.NotEmpty(t, use.Username)
	require.NotEmpty(t, use.PasswordHash)
	require.NotZero(t, use.ID)
	// Now we can get the user by username
	username := fmt.Sprintf("John@123_TestGetUserByUsername_%d", timestamp)
	use, err = testQueries.GetUserByUsername(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, use)

	// Clean up: Delete the user
	err = testQueries.DeleteUser(context.Background(), use.ID)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), 0)
	require.NoError(t, err)
}
