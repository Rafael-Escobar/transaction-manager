package postgrees

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/transaction-manager/internal/config"
)

type Client struct {
	db *sqlx.DB
}

// NewClient creates a new Postgres client.
func NewClient(dsn string, connectionConfig config.RelationalDBConnection) (*Client, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sqlx failed to connect to db: %w", err)
	}

	db.SetMaxIdleConns(connectionConfig.MaxIdleConns)
	db.SetMaxOpenConns(connectionConfig.MaxOpenConns)
	db.SetConnMaxLifetime(connectionConfig.MaxLifeTime)
	db.SetConnMaxIdleTime(connectionConfig.MaxIdleTime)

	return &Client{db: db}, nil
}
