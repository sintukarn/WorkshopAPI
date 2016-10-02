package services

import (
	"model"
	"dao"
	"time"
	"gopkg.in/mgo.v2/bson"
)
var (
	findoneDeactiveDevBefore = dao.FindOne
	findoneDeactiveDevAfter = dao.FindOne
	updateDeactiveDev = dao.Update
)

func DeactiveDevService(name string,b string)(*model.Dev,error){
	var resultDev model.Dev
	err := findoneDeactiveDevBefore("dev", name, &resultDev)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"update_date": update_date, "active": b}}
	err = updateDeactiveDev("dev", find, update)
	if err != nil {
		return nil, err
	}
	var result model.Dev
	err = findoneDeactiveDevAfter("dev", name, &result)
	if err != nil {
		return nil, err
	}
	return &result,nil
}
