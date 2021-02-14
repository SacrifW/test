package db

import (
	"../models"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

type dataBaseModule struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func GetBlock(blockNumber int) (models.Block, error) {
	var block models.Block

	var session, err = mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Println("db.go -> GetBlock() -> mgo.Dial -> error-db:", err)
		return block, err
	}

	var dbModule = &dataBaseModule{
		session:    session,
		collection: session.DB("BlockDB").C("blocks"),
	}

	blockNumberStr := fmt.Sprintf("0x%x", blockNumber)
	filter := models.Obj{"result.transactions.blockNumber": blockNumberStr}

	if err := dbModule.collection.Find(filter).One(&block); err != nil {
		log.Println("db.go -> GetBlock() -> Find() -> error-db:", err)
		return block, err
	}

	return block, nil
}
