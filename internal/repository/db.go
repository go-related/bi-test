package repository

import (
	"entdemo/ent"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	client *ent.Client
}

func NewDb(conn string) (*DB, error) {
	client, err := ent.Open("postgres", conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed opening connection to postgres")
	}

	return &DB{client: client}, nil
}
