package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"squid/postgres-stress-test/http"
	"squid/postgres-stress-test/logger"
	"squid/postgres-stress-test/postgres"
	"syscall"

	_ "github.com/lib/pq"
)

var (
	apiAddress       = ":8888"
	postgresUser     = "postgres"
	postgresPassword = ""
	postgresHost     = "localhost"
	//postgresPassword = ""
	//postgresHost     = "squid-aurora-postgresql-instance-1.cn72nfyrfryl.eu-north-1.rds.amazonaws.com"
	postgresPort = "5432"
	postgresDB   = "postgres"
	postgresSSL  = "disable"
)

type logRepository struct {
	db *sql.DB
}

func main() {
	// Postgres setup.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort,
		postgresDB,
		postgresSSL,
	)
	fmt.Println("connecting to ", dsn)
	// Postgres...
	db, err := sql.Open("postgres", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println("couldn't connect to db")
		os.Exit(1)
	}

	db.SetMaxOpenConns(800)
	db.SetMaxIdleConns(750)

	err = db.Ping()
	if err != nil {
		fmt.Println("db ping failed")
		os.Exit(1)
	}

	incomingRepo := postgres.NewIncomingRepository(db)
	logService := logger.NewService(incomingRepo)

	errorChannel := make(chan error)

	go func() {
		httpServer := http.NewServer(apiAddress, db, logService)
		errorChannel <- httpServer.Open()
	}()
	// Capture interrupts.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-c)
	}()

	// Wait (indefinitely) for any error.
	if err := <-errorChannel; err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
