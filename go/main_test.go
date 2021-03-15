package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	_ "testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestCredFromJSON(t *testing.T) {

	json := []byte(`{"username": "a","password": "b","wallet": "wallet-1"}`)
	user_cred := FromJSON(json)
	assert.Equal(t, UserCredentials{Username: "a", Password: "b", Wallet: "wallet-1"}, user_cred, "User cred JSON unmarshalling wrong.")
}

func TestResponses(t *testing.T) {
	requestDebit, err := json.Marshal(map[string]string{
		"username":    "a",
		"password":    "b",
		"wallet":      "wallet-1",
		"transaction": "10",
	})

	requestCredit, err := json.Marshal(map[string]string{
		"username":    "a",
		"password":    "b",
		"wallet":      "wallet-1",
		"transaction": "-10",
	})

	requestWrong, err := json.Marshal(map[string]string{
		"username":    "a",
		"password":    "b",
		"wallet":      "waaaallet-1",
		"transaction": "10",
	})

	requestOvercharge, err := json.Marshal(map[string]string{
		"username":    "a",
		"password":    "b",
		"wallet":      "wallet-1",
		"transaction": "100000000000000",
	})

	request, err := http.NewRequest("POST", "http://localhost:8080/manage-balance/input", bytes.NewBuffer(requestDebit))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	request, err = http.NewRequest("POST", "http://localhost:8080/balance/input", bytes.NewBuffer(requestCredit))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	request, err = http.NewRequest("POST", "http://localhost:8080/manage-balance/input", bytes.NewBuffer(requestCredit))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	request, err = http.NewRequest("POST", "http://localhost:8080/credit/input", bytes.NewBuffer(requestCredit))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	request, err = http.NewRequest("POST", "http://localhost:8080/debit/input", bytes.NewBuffer(requestCredit))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	request, err = http.NewRequest("POST", "http://localhost:8080/manage-balance/input", bytes.NewBuffer(requestWrong))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	request, err = http.NewRequest("POST", "http://localhost:8080/manage-balance/input", bytes.NewBuffer(requestOvercharge))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

}

// POST http://localhost:8080/manage-balance/input

// {"username":"a","password":"b","wallet":"wallet-1", "transaction":10}
// {"username":"a","password":"b","wallet":"wallet-1"}
