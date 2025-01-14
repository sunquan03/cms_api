package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"strconv"

	"log"
	"os"
)

type PostgresLayer struct {
	db *sqlx.DB
}

func NewPostgresLayer(db *sqlx.DB) *PostgresLayer {
	return &PostgresLayer{db: db}
}

func NewDB() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %s", err.Error())
	}

	dbName := os.Getenv("pg_db")
	dbHost := os.Getenv("pg_host")
	dbPort := os.Getenv("pg_port")
	dbUser := os.Getenv("pg_user")
	dbPassword := os.Getenv("pg_password")
	maxOpenConns := os.Getenv("pg_max_open_conns")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	connConfig, err := pgx.ParseConfig(dbUri)
	if err != nil {
		log.Fatal(err)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)
	conn, err := sqlx.Open("postgres", connStr)

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	maxConns, err := strconv.Atoi(maxOpenConns)
	if err != nil {
		conn.SetMaxOpenConns(5)
	}

	conn.SetMaxIdleConns(maxConns)

	log.Println("established connection to postgres")
	return conn
}
