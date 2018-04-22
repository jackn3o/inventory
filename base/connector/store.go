package connector

import (
	"../configuration"
	"github.com/garyburd/redigo/redis"
	mgo "gopkg.in/mgo.v2"
)

// Store is db connection
type Store struct {
	DB *mgo.Session
	KV *redis.Pool
}

// New create and return mongo and redis connection
func New(config configuration.Config) *Store {
	db := createMongoConnection(config)

	return &Store{
		DB: db,
	}
}
