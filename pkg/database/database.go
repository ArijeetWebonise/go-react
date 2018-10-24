package database

import (
	"database/sql"

	"github.com/Arijeet-webonise/go-react/app/config"
	_ "github.com/lib/pq"
)

// DatabaseConnectionInitialiser encapsulates DB object
type DatabaseConnectionInitialiser interface {
	Initialise(map[string]string) (*sql.DB, error)
}

// DatabaseConfig wrapper for DB Config
type DatabaseWrapper struct {
	DB *sql.DB
}

// InitialiseConnection init DB
func (dw *DatabaseWrapper) Initialise(dbConfig *config.DbConfig) (*sql.DB, error) {
	db, dbConnErr := sql.Open(dbConfig.DbDriverName, dbConfig.DbDataSource)
	return db, dbConnErr
}
