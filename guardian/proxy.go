package main

import (
	"fmt"
	"sql-sentry/client"
)

type Proxy struct {
	conn       *client.Conn
	allowedSQL *regex.Regexp
}

func (p Proxy) UseDB(dbName string) error {
	return nil
}

func (p Proxy) HandleQuery(query string) (*Result, error) {
	return nil, fmt.Errorf("not supported now")
}

func (p Proxy) HandleFieldList(table string, fieldWildcard string) ([]*Field, error) {
	return nil, fmt.Errorf("not supported now")
}
func (p Proxy) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	return 0, 0, nil, fmt.Errorf("not supported now")
}
func (p Proxy) HandleStmtExecute(context interface{}, query string, args []interface{}) (*Result, error) {
	return nil, fmt.Errorf("not supported now")
}

func (p Proxy) HandleStmtClose(context interface{}) error {
	return nil
}
