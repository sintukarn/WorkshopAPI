package services

import (
	"dao"
	"model"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func InsertSquadToBUService(unit string,target string)(model.BU,error) {
	var resultBU model.BU
	err := dao.FindOne("BU",target,&resultBU)
	if err!= nil {
		return model.BU{},err
	}
	var resultSquad model.Squad
	err = dao.FindOne("squad",unit,&resultSquad)
	if err!= nil{
		return model.BU{},err
	}
	update_date := time.Now().Unix()
	find := bson.M{"name":resultBU.Name}
	update := bson.M{"$set":bson.M{"squads":append(resultBU.Squads,resultSquad),"update_date":update_date}}
	err = dao.Update("BU",find,update)
	if err != nil {
		return model.BU{},err
	}
	var result model.BU
	err = dao.FindOne("BU",target,&result)
	return result,err
}
