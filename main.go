package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var clients = make(map[net.Conn]string)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Get username
	fmt.Fprintln(conn, "Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	clients[conn] = name

	broadcast(fmt.Sprintf("%s joined the chat!", name), conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		broadcast(fmt.Sprintf("%s: %s", name, msg), conn)
	}

	broadcast(fmt.Sprintf("%s left the chat", name), conn)
	delete(clients, conn)
}
func broadcast(message string, sender net.Conn) {
	for conn := range clients {
		if conn != sender {
			fmt.Fprintln(conn, message)
		}
	}
	fmt.Println(message) // server console
}
