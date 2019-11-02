package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s", msg)
	return func() {
		log.Printf("Exit %s (%.3gs)", msg, time.Since(start).Seconds())
	}
}

func bigSlowOperation() {
	defer trace("big slow operation")()

	time.Sleep(10 * time.Second)
}

func main() {
	bigSlowOperation()
}
