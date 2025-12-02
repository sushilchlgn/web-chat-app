package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			msg, _ := serverReader.ReadString('\n')
			fmt.Print(msg)
		}
	}()

	consoleReader := bufio.NewReader(os.Stdin)
	for {
		text, _ := consoleReader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Fprintln(conn, text)
		if strings.ToLower(text) == "exit" {
			break
		}
	}
}
