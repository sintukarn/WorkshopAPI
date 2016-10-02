package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"time"
)

func DeactiveSquadService(name string, b string) (*model.Squad, error) {
	var resultSquad model.Squad
	err := dao.FindOne("squad", name, &resultSquad)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	var resultsquad model.Squad
	err = dao.FindOne("squad", name, &resultsquad)
	if err != nil {
		return nil, err
	}
	var newDevs []*model.Dev
	for i := range resultsquad.Devs {
		dev,err := DeactiveDevService(resultsquad.Devs[i].Name,b)
		if err != nil {
			return nil, err
		}
		newDevs = append(newDevs,dev)
	}
	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"devs" : newDevs ,"update_date": update_date, "active": b}}
	err = dao.Update("squad", find, update)
	if err != nil {
		return nil, err
	}
	var result model.Squad
	err = dao.FindOne("squad", name, &result)
	return &result, err
}
