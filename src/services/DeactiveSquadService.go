package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"time"
)

func DeactiveSquadService(name string, b string) (interface{}, error) {
	var resultSquad model.Squad
	err := dao.FindOne("squad", name, &resultSquad)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"update_date": update_date, "active": b}}
	err = dao.Update("squad", find, update)
	if err != nil {
		return nil, err
	}
	var resultsquad model.Squad
	err = dao.FindOne("squad", name, &resultsquad)
	if err != nil {
		return nil, err
	}
	for i := range resultsquad.Devs {
		dev := bson.M{"name": resultsquad.Devs[i].Name}
		update = bson.M{"$set": bson.M{"update_date": update_date, "active": b}}
		err = dao.Update("dev", dev, update)
		if err != nil {
			return nil, err
		}
	}
	var result model.Squad
	err = dao.FindOne("squad", name, &result)
	return &result, err
}
