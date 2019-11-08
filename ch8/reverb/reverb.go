package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Println(shout, delay)
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for firstEcho := true; input.Scan(); firstEcho = false {
		wg.Add(1)
		go func(shout string) {
			defer wg.Done()
			echo(c, shout, time.Second)
		}(input.Text())

		if firstEcho {
			go func() {
				fmt.Println("Start Waiting")
				wg.Wait()
				fmt.Println("Start Closing")
				if cwc, ok := c.(*net.TCPConn); ok {
					cwc.CloseWrite()
				}
			}()
		}
	}
}

func handleConn2(c net.Conn) {
	in := make(chan string)

	go func(c net.Conn) {
		input := bufio.NewScanner(c)
		for input.Scan() {
			in <- input.Text()
		}
		fmt.Println("Am I overflowed? Kidding me!")
	}(c)

	tick := time.Tick(time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			// do nothing
		case shout := <-in:
			countdown = 10
			echo(c, shout, time.Second)
		}
	}

	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn2(conn)
	}
}
