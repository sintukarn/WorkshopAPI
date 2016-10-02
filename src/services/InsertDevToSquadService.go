package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"log"
)

var (
	findoneInsertDevtoSquadBefore = dao.FindOne
	findoneInsertDevtoSquadAfter  = dao.FindOne
	updateInsertDevtoSquad        = dao.Update
)

func InsertDevToSquadService(unit string, target string) (*model.Squad, error) {
	var resultSquad model.Squad
	log.Println(target)
	err := findoneInsertDevtoSquadBefore("squad", target, &resultSquad)
	if err != nil {
		return nil, err
	}
	log.Println(resultSquad)
	var resultDev model.Dev
	err = findoneInsertDevtoSquadBefore("dev", unit, &resultDev)
	if err != nil {
		return nil, err
	}
	update_date := update_date_func()
	find := bson.M{"name": resultSquad.Name}
	update := bson.M{"$set": bson.M{"devs": append(resultSquad.Devs, resultDev), "update_date": update_date}}
	err = updateInsertDevtoSquad("squad", find, update)
	if err != nil {
		log.Println("not found ?")
		return nil, err
	}
	var result model.Squad
	err = findoneInsertDevtoSquadAfter("squad", target, &result)
	return &result, err
}
