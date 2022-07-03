package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	DriverName string = "postgres"
)

type Config struct {
	// [ConnectionString](https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters)
	ConnectionString string
}

func (cfg *Config) beginTx(ctx context.Context) (*sql.Tx, func() error, error) {
	db, err := sql.Open(DriverName, cfg.ConnectionString)
	if err != nil {
		return nil, nil, err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		db.Close()
		return nil, nil, err
	}

	return tx, db.Close, err
}
