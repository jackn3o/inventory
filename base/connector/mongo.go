package connector

import (
	"time"

	"../configuration"
	mgo "gopkg.in/mgo.v2"
)

// createMongoConnection
func createMongoConnection(config configuration.Config) *mgo.Session {
	database := config.GetString(configuration.DatabaseName)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{config.GetString(configuration.DatabaseHost)},
		Timeout:  time.Duration(config.GetInt(configuration.DatabaseTimeout)) * time.Second,
		Username: config.GetString(configuration.DatabaseUsername),
		Password: config.GetString(configuration.DatabasePassword),
		Database: database,
	})

	if err != nil {
		panic(err)
	}

	return session
}
