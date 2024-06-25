package main

import (
 "errors"
 "fmt"
)

type BankAccount struct {
 Balance float64
 Owner   string
}

func (a *BankAccount) Deposit(amount float64) {
 a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
 if amount > a.Balance {
  return errors.New("Недостаточно средств")
 }
 a.Balance -= amount
 return nil
}

func (a *BankAccount) GetBalance() float64 {
 return a.Balance
}

func main() {
 account := BankAccount{Balance: 5000, Owner: "Behruz"}

 account.Deposit(20)

 err := account.Withdraw(50)
 if err != nil {
  fmt.Println(err)
 }

 balance := account.GetBalance()
 fmt.Println(balance)
}