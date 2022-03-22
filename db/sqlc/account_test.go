package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rtpa25/go_api_worflow/utils"
	"github.com/stretchr/testify/assert"
)

func createRadnomAccount(t *testing.T) Account {
	user := createRadnomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RadomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotNil(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Currency, account.Currency)

	assert.NotNil(t, account.ID)
	assert.NotNil(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRadnomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRadnomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	assert.NoError(t, err)

	assert.Equal(t, acc1.ID, acc2.ID)
	assert.Equal(t, acc1.Balance, acc2.Balance)
	assert.Equal(t, acc1.Currency, acc2.Currency)
	assert.Equal(t, acc1.Owner, acc2.Owner)
	assert.Equal(t, acc1.CreatedAt, acc2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRadnomAccount(t)
	args := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: utils.RandomMoney(),
	}

	acc2, err := testQueries.UpdateAccount(context.Background(), args)
	assert.NoError(t, err)
	assert.Equal(t, acc1.ID, acc2.ID)
	assert.Equal(t, acc1.Currency, acc2.Currency)
	assert.Equal(t, acc1.CreatedAt, acc2.CreatedAt)
	assert.Equal(t, acc1.Owner, acc2.Owner)
	assert.Equal(t, args.Balance, acc2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRadnomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	assert.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	if assert.Error(t, err) {
		assert.Equal(t, sql.ErrNoRows.Error(), err.Error())
	}
	assert.Empty(t, acc2)
}

func TestListAllAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRadnomAccount(t)
	}

	//skip first 5 and return next 5
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(accounts))

	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}
