package main

import (
	"fmt"
	"net"

	"github.com/siddontang/go-mysql/server"
)

type Behavior int

const (
	DENY = iota
	WARN = iota
	ALL  = iota
)

type Config struct {
	DatabasePattern string
	AllowedSQLFile  string
	ProxyPort       int
	DB              struct {
		User string
		Pass string
		Host string
		Port int
	}
	ProxyBehavior Behavior
}

func loadConfig(configFile string) (Config, error) {
	return nil, nil
}

func main() {

	//TODO: parse cmd line args

	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading config file %s: %s", configFile, err.Error())
	}

	l, _ := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))

	c, _ := l.Accept()

	proxy := Proxy{
		conn:       nil,
		allowedSQL: sqlPattern,
		behavior:   config.Behavior,
	}

	// Create a connection with user root and an empty password
	// We only an empty handler to handle command too
	conn, _ := server.NewConn(c, config.User, config.Pass, proxy)

	for {
		conn.HandleCommand()
	}

}
