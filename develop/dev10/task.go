package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigCh
		fmt.Println("Close signal")
		os.Exit(0)
	}()

	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr())

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data := scanner.Bytes()
			_, err := conn.Write(data)
			if err != nil {
				fmt.Printf("Error while wrtiting to socket: %s\n", err.Error())
				break
			}

			response := make([]byte, 1024)
			_, err = conn.Read(response)
			if err != nil {
				fmt.Printf("Error while reading from socket: %s\n", err.Error())
				break
			}

			fmt.Println(string(response))
		}
	}()

	<-sigCh
	fmt.Println("Connection closed")
}
