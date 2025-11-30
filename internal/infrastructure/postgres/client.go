package postgres

import (
	"fmt"
	"geoserv/internal/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	DB *sqlx.DB
}

func NewPostgresClient(cfg config.PostgresConfig) *PostgresClient {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.SslMode))
	if err != nil {
		log.Panicf("cannot open database: %v", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Panicf("database not pinging: %v", err)
		return nil
	}
	createTables(db)
	return &PostgresClient{DB: db}
}

func createTables(db *sqlx.DB) error {
	// create drivers table
	if err := execQuery(db, `create table if not exists drivers (
		id serial primary key,
		name varchar(50),
		is_busy boolean
	);`); err != nil {
		return err
	}
	// create drivers positions table
	if err := execQuery(db, `create table if not exists driver_positions (
		id serial primary key,
		driver_id text not null,
		lat double precision,
		lon double precision,
		timestamp timestamptz
	);`); err != nil {
		return err
	}
	// create zones table
	if err := execQuery(db, `create table if not exists zones (
		id serial primary key,
		workload double precision	
	);`); err != nil {
		return err
	}
	return nil
}

func execQuery(db *sqlx.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
