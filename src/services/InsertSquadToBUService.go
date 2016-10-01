package services

import (
	"dao"
	"model"
	"time"
	"gopkg.in/mgo.v2/bson"
)
var(
	findoneInsertSquadtoBUBefore = dao.FindOne
	findoneInsertSquadtoBUAfter = dao.FindOne
	updateInsertSquadtoBU = dao.Update
)

func InsertSquadToBUService(unit string,target string)(*model.BU,error) {
	var resultBU model.BU
	err := findoneInsertSquadtoBUBefore("BU",target,&resultBU)
	if err!= nil {
		return nil,err
	}
	var resultSquad model.Squad
	err = findoneInsertSquadtoBUBefore("squad",unit,&resultSquad)
	if err!= nil{
		return nil,err
	}
	update_date := time.Now().Unix()
	find := bson.M{"name":resultBU.Name}
	update := bson.M{"$set":bson.M{"squads":append(resultBU.Squads,resultSquad),"update_date":update_date}}
	err = updateInsertSquadtoBU("BU",find,update)
	if err != nil {
		return nil,err
	}
	var result model.BU
	err = findoneInsertSquadtoBUAfter("BU",target,&result)
	return &result,err
}
