package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/guntoroyk/go-user-api/config"
)

const (
	dbCredentialsFormat = "%s:%s@tcp(%s:%s)/%s"
)

// NewDB will create new DB connection
func NewDB(cfg *config.DBConfig) *sql.DB {
	address := fmt.Sprintf(dbCredentialsFormat,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	fmt.Println("[Database] connecting to DB: " + address)
	db, err := sql.Open("mysql", address)

	if err != nil {
		log.Fatal("[Database] failed connecting to DB: " + address + ", err: " + err.Error())
	}

	// check established connection with DB
	if err := db.Ping(); err != nil {
		log.Fatal("[Database] db is unreachable: " + address + ", err: " + err.Error())
	}

	return db
}
