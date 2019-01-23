package repository

import (
	"github.com/juju/mgosession"
)

// MongoDB defines the pool and database name for a MongoDB connection.
type MongoDB struct {
	pool *mgosession.Pool
	db   string
}

// New creates a new Repository layer sharing the MongoDB connection.
func New(p *mgosession.Pool, db string) *MongoDB {
	return &MongoDB{
		pool: p,
		db:   db,
	}
}
