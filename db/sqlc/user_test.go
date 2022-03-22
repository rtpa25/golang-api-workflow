package db

import (
	"context"
	"testing"

	"github.com/rtpa25/go_api_worflow/utils"
	"github.com/stretchr/testify/assert"
)

func createRadnomUser(t *testing.T) User {
	hashedPass, err := utils.HashPassword(utils.RandomString(6))
	assert.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RadnomOwner(),
		HashedPassword: hashedPass,
		FullName:       utils.RadnomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, arg.Email, user.Email)
	assert.Equal(t, arg.FullName, user.FullName)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.True(t, user.PasswordChangedAt.IsZero())
	assert.NotNil(t, user.CreatedAt)

	return user
}

func TestCreateUsert(t *testing.T) {
	createRadnomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRadnomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	assert.NoError(t, err)

	assert.Equal(t, user1.Email, user2.Email)
	assert.Equal(t, user1.FullName, user2.FullName)
	assert.Equal(t, user1.HashedPassword, user2.HashedPassword)
	assert.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	assert.Equal(t, user1.CreatedAt, user2.CreatedAt)
}
