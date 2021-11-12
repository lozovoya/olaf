package main

import (
	"log"
	"net"
	"os"
)

const (
	defaultPort     = "9999"
	defaultHost     = "0.0.0.0"
	defaultDSN      = "postgres://app:pass@olafdb:5432/olafdb"
)

func main() {
	port, ok := os.LookupEnv("MARKET_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("MARKET_HOST")
	if !ok {
		host = defaultHost
	}

	dsn, ok := os.LookupEnv("MARKET_DB")
	if !ok {
		dsn = defaultDSN
	}

	if err := execute(net.JoinHostPort(host, port), dsn); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute (addr string, dsn string) (err error) {

	return nil
}