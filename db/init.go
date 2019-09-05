package db

import (
	"time"

	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

var SessionOrginal *mgo.Session
var databaseName string

var (
	MongoDBHosts string
	AuthDatabase string
	AuthUserName string
	AuthPassword string
)

func nativeConnection() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mgoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	log.Info("Mongodb connected")
	mgoSession.SetMode(mgo.Monotonic, true)
	SessionOrginal = mgoSession
}

func init() {
	MongoDBHosts = "127.0.0.1"
	AuthDatabase = "admin"
	AuthUserName = "mongoadmin"
	AuthPassword = "secret"
	databaseName = "go3008"

	nativeConnection()
}
