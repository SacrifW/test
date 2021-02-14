package db

import (
	"../models"
	"gopkg.in/mgo.v2"
	"log"
)

type dataBaseModule struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func GetBlock(blockNumber uint64) (models.Block, error) {
	var block models.Block

	var session, err = mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Println("db.go -> GetBlock() -> mgo.Dial -> error-db", err)
		return block, err
	}

	var dbModule = &dataBaseModule{
		session:    session,
		collection: session.DB("blockDB").C("blocks"),
	}

	filter := models.Obj{"result.transactions.blockNumber": blockNumber}

	if err := dbModule.collection.Find(filter).One(&block); err != nil {
		log.Println("db.go -> GetBlock() -> Find() -> error-db", err)
		return block, err
	}

	return block, nil
}
