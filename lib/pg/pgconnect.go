package pg

import (

	"github.com/sirupsen/logrus"

	"fmt"

	"github.com/jmoiron/sqlx"

)

// Replace with your actual values or load from constants/env
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "postgress_db_test"
	sslmode  = "disable"
)

func NewLocalPGInstance() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logrus.Errorf("failed to open db connection: %v", err)
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	// Ping to confirm connection is alive
	if err := db.Ping(); err != nil {
		logrus.Errorf("failed to ping db: %v", err)
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Optional tuning
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(5)

	return db, nil
}
