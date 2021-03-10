package main

import (
	"testing"
	_ "testing"

	_ "github.com/stretchr/testify/assert"
)

func TestCredFromJSON(t *testing.T) {
	json := []byte(`{"username": "a","password": "b","wallet": "wallet-1"}`)
	user_cred := FromJSON(json)
	assert.Equal(t, UserCredentials{Username: "a", Password: "b", Wallet: "wallet-1"}, user_cred, "User cred JSON unmarshalling wrong.")
}
