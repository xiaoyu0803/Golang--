package main

import "chat/server"

func main() {
	server := server.NewServer("127.0.0.1", 8888)
	server.Start()
}
