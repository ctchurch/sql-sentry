package main

import (
	"net"

	"github.com/siddontang/go-mysql/server"
)

func main() {

	l, _ := net.Listen("tcp", "127.0.0.1:4000")

	c, _ := l.Accept()

	// Create a connection with user root and an empty password
	// We only an empty handler to handle command too
	conn, _ := server.NewConn(c, config.User, config.Pass, proxy)

	for {
		conn.HandleCommand()
	}

}
