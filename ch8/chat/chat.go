package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	signin   = make(chan client)
	signout  = make(chan client)
	messages = make(chan string)
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-signin:
			clients[cli] = true
		case cli := <-signout:
			delete(clients, cli)
			close(cli)
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		}
	}
}

func messageWriter(c net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintf(c, "%s\n", msg)
	}
}

func handleConn(c net.Conn) {
	ch := make(chan string)
	go messageWriter(c, ch)

	who := c.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + " has arrived"
	signin <- ch

	input := bufio.NewScanner(c)
	for input.Scan() {
		messages <- input.Text()
	}

	signout <- ch
	messages <- who + " has left"
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
