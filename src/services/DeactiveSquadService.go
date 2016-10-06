package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"time"
)

var (
	findoneDeactiveSquadBefore = dao.FindOne
	findoneDeactiveSquadAfter  = dao.FindOne
	updateDeactiveSquad        = dao.Update
	deacDev                    = DeactiveDevService
)

func DeactiveSquadService(name string, b string) (*model.Squad, error) {
	var resultSquad model.Squad
	err := findoneDeactiveSquadBefore("squad", name, &resultSquad)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	var newDevs []*model.Dev
	for i := range resultSquad.Devs {
		dev, err := deacDev(resultSquad.Devs[i].Name, b)
		if err != nil {
			return nil, err
		}
		newDevs = append(newDevs, dev)

	}

	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"devs": newDevs, "update_date": update_date, "active": b}}
	err = updateDeactiveSquad("squad", find, update)
	if err != nil {
		return nil, err
	}
	UpdateSquadinBUService(name)
	var result model.Squad
	err = findoneDeactiveSquadAfter("squad", name, &result)
	return &result, err
}
