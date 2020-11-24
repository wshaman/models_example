package model

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // init PGSQL driver

)

var db *sqlx.DB

// Model interface for all models
type Model interface {
	UserModel
}

type model struct{}

//@nb: this is questionable solution. Set here as an all-in-one drop-in usage example.
func init() {
	database := os.Getenv("PGSQL_DATABASE")
	if database == "" {
		database = "postgres"
	}

	hostname := os.Getenv("PGSQL_HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}

	username := os.Getenv("PGSQL_USERNAME")
	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("PGSQL_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	port := os.Getenv("PGSQL_PORT")
	if port == "" {
		port = "5432"
	}

	sslMode := os.Getenv("PGSQL_SSLMODE")
	if sslMode == "" {
		sslMode = "require"
	} else {
		/* https://godoc.org/github.com/lib/pq
		Valid values for sslmode are:
		* disable - No SSL
		* require - Always SSL (skip verification)
		* verify-ca - Always SSL (verify that the certificate presented by the server was
		  signed by a trusted CA)
		* verify-full - Always SSL (verify that the certification presented by the server was signed
		  by a trusted CA and the server host name matches the one in the certificate)
		*/
		if sslMode != "disable" && sslMode != "require" {
			log.Fatalf("Incorrect value '%s' for postgres sslmode", sslMode)
		}
	}

	dataSource := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		database, hostname, port, username, password, sslMode)
	var err error
	db, err = sqlx.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
}

// NewModel returns model
func NewModel() (Model, error) {
	return &model{}, nil
}

// NewModelWithMetrics returns model wrapped with metrics
func NewModelWithMetrics() (Model, error) {
	m, err := NewModel()
    if err != nil {
        return nil, err
    }
    return newMetrics(NewModel())
}

// TryTx executes a function within a single transaction, rolling back on any returned error.
func TryTx(f func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
	return err
}

type dbDates struct {
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at"`
}
