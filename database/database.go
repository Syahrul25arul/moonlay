package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"moonlay/helper"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5431", "moonlay_test")
	db, err := sql.Open("postgres", connStr)
	helper.PanicIFError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func TruncateAllTable(db *sql.DB) {
	sql := "TRUNCATE TABLE transactions,products,customers restart identity"
	db.ExecContext(context.Background(), sql)
}
