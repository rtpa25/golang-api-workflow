package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashingOfPassword(t *testing.T) {
	password := RandomString(6)
	hashedPass1, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPass1)
	assert.NoError(t, CheckPassword(password, hashedPass1))

	wrongPassword := RandomString(6)
	assert.EqualError(t, CheckPassword(wrongPassword, hashedPass1), bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPass2, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPass2)
	assert.NotEqual(t, hashedPass1, hashedPass2)
}
