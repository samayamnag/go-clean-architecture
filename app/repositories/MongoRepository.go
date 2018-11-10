package repositories

import (
	"github.com/juju/mgosession"
)

const (
	USER_COLLECTION = "users"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

//NewMongoRepository create new repository
func NewMongoRepository(p *mgosession.Pool, db string) *MongoRepository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}
