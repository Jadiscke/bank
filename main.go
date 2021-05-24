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

func (c *CheckingAccount) Withdraw(value float64) string {
	canWithdraw := value <= c.balance && value > 0.0

	if canWithdraw {
		c.balance -= value
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func main() {
	checkingAccount := createAccount("Silvia")

	fmt.Println(checkingAccount.Withdraw(-200))

	fmt.Println(checkingAccount.balance)
}
