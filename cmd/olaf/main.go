package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"net"
	"net/http"
	"olaf/internal/api/httpserver"
	v1 "olaf/internal/api/v1"
	"olaf/internal/repository"
	"os"
)

const (
	defaultPort = "9999"
	defaultHost = "0.0.0.0"
	defaultDSN  = "postgres://app:pass@olafdb:5432/olafdb"
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
		os.Exit(1)
	}
}

func execute(addr string, dsn string) (err error) {

	lg := zap.NewExample()
	defer lg.Sync()

	usersCtx := context.Background()
	usersPool, err := pgxpool.Connect(usersCtx, dsn)
	if err != nil {
		lg.Error("execute ", zap.Error(err))
		return err
	}
	usersRepo := repository.NewUsersPool(usersPool)
	usersController := v1.NewUsers(usersRepo, lg)

	groupsCtx := context.Background()
	groupsPool, err := pgxpool.Connect(groupsCtx, dsn)
	if err != nil {
		lg.Error("execute ", zap.Error(err))
		return err
	}
	groupsRepo := repository.NewGroupsRepo(groupsPool)
	groupsController := v1.NewGroups(groupsRepo, usersRepo, lg)

	router := httpserver.NewRouter(chi.NewRouter(), usersController, groupsController, lg)
	server := http.Server{
		Addr:    addr,
		Handler: &router,
	}
	return server.ListenAndServe()
}
