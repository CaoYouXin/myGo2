package main

import (
	"fmt"
	"time"
)

var (
	deposits = make(chan int)
	balances = make(chan int)
)

// Deposit xxx
func Deposit(amount int) { deposits <- amount }

// Balance xxx
func Balance() int { return <-balances }

func teller() {
	fmt.Println("start a teller")
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Println("send ", balance)
		}
	}
}

func init() {
	go teller()
}

func main() {
	fmt.Println(Balance())
	Deposit(100)
	Deposit(200)
	Deposit(300)
	time.Sleep(time.Second)
	fmt.Println(Balance())
	time.Sleep(time.Second)
}
