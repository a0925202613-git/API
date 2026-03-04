package database

import (
	"database/sql"

	"go-api-practice/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const createTablesSQL = `
CREATE TABLE IF NOT EXISTS items (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
`

// Connect 建立 PostgreSQL 連線並建立資料表
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", config.DatabaseURL())
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	_, err = DB.Exec(createTablesSQL)
	return err
}
