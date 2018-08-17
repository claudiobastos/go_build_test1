package main

import "fmt"

//Connection Dados de conexao
type Connection struct {
	host string
	user string
	pass string
	base string
	port uint64
}

// return Data Source Name for database connection
// e.g.: "user:password@tcp(host:port)/schema"
func (c Connection) getDSN() string {
	//"user:password@tcp(host:port)/schema"
	//return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.user, c.pass, c.host, c.port, c.base)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.user, c.pass, c.host, c.base)
}

// return Data Source Name for database connection
// e.g.: "user:password@tcp(host:port)/schema"
func (c Connection) getUndisclosedDSN() string {
	//"user:password@tcp(host:port)/schema"
	return fmt.Sprintf("%s:*****@tcp(%s:%d)/%s", c.user, c.host, c.port, c.base)
}
