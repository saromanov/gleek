// Package storage contains handling with db(Postgesql)
package storage

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saromanov/gleek/config"
)

var (
	errNoConfig = errors.New("config is not defined")
	errNoCreds  = errors.New("name, password or user is not defined for storage")
)

// Storage implements db handling with Postgesql
type Storage struct {
	db *sqlx.DB
}

// New provides init for postgesql storage
func New(s *config.Config) (*Storage, error) {
	if s == nil {
		return nil, errNoConfig
	}
	if s.Name == "" || s.Password == "" || s.User == "" {
		return nil, errNoCreds
	}
	args := "dbname=gleek2"
	if s.Name != "" && s.Password != "" && s.User != "" {
		args += fmt.Sprintf(" user=%s dbname=%s password=%s", s.User, s.Name, s.Password)
	}
	db, err := sqlx.Connect("postgres", args)
	if err != nil {
		log.Fatalln(err)
	}
	return &Storage{
		db: db,
	}, nil
}

// Close provides closing of db
func (s *Storage) Close() error {
	return s.db.Close()
}
