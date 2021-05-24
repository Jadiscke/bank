package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type CheckingAccount struct {
	owner         string
	agencyNumber  string
	accountNumber string
	balance       float64
}

func createAccount(name string) CheckingAccount {
	maxAgencyValue := big.NewInt(999)
	maxAccountValue := big.NewInt(999999)

	randomAgencyValue, _ := rand.Int(rand.Reader, maxAgencyValue)
	randomAccountValue, _ := rand.Int(rand.Reader, maxAccountValue)
	return CheckingAccount{
		owner:         name,
		agencyNumber:  strconv.FormatInt(randomAgencyValue.Int64(), 10),
		accountNumber: strconv.FormatInt(randomAccountValue.Int64(), 10),
		balance:       0,
	}
}

func main() {
	checkingAccount := createAccount("Vinicius")

	fmt.Println(checkingAccount)
}
