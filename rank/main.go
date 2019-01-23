package main

import (
	"log"

	"github.com/juju/mgosession"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/controller"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/middlewares/config"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/middlewares/routing"
	"github.com/laurenmariaferreira/clean-architecture-case-study/rank/repository"
	mgo "gopkg.in/mgo.v2"
)

const port = ":8899"

func main() {
	Rank()
}

// Rank starts the routine for Rank's app.
func Rank() {
	session, err := mgo.Dial(config.MongoDBHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MongoDBPool)
	defer pool.Close()

	repo := repository.New(pool, config.MongoDBDatabaseName)

	controllers := controller.New(repo)

	router := routing.Router(controllers)

	router.Run(config.Port)
	log.Printf("Running router on port %s", config.Port)
}
