package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"time"
)

var (
	fineoneDeactiveBUBefore = dao.FindOne
	fineoneDeactiveBUAfter  = dao.FindOne
	updateDeactiveBU        = dao.Update
	deacSquad               = DeactiveSquadService
)

func DeactiveBUService(name string, b string) (*model.BU, error) {
	var resultBU model.BU
	err := fineoneDeactiveBUBefore("BU", name, &resultBU)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	var newSquads []*model.Squad
	for i := range resultBU.Squads {
		squad, err := deacSquad(resultBU.Squads[i].Name, b)
		if err != nil {
			return nil, err
		}
		newSquads = append(newSquads, squad)
	}
	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"squads": newSquads, "update_date": update_date, "active": b}}
	err = updateDeactiveBU("BU", find, update)
	if err != nil {
		return nil, err
	}
	var findBU model.BU
	err = fineoneDeactiveBUAfter("BU", name, &findBU)
	return &findBU, err
}
