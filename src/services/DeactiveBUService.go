package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
	"time"
)

func DeactiveBUService(name string, b string) (interface{}, error) {
	var resultBU model.BU
	err := dao.FindOne("BU", name, &resultBU)
	if err != nil {
		return nil, err
	}
	update_date := time.Now().Unix()
	find := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"update_date": update_date, "active": b}}
	err = dao.Update("BU", find, update)
	if err != nil {
		return nil, err
	}
	var resultbu model.BU
	err = dao.FindOne("BU", name, &resultbu)

	for i := range resultbu.Squads {
		squad := bson.M{"name": resultbu.Squads[i].Name}
		update = bson.M{"$set": bson.M{"update_date": update_date}, "active": b}
		err = dao.Update("squad", squad, update)
		if err != nil {
			return nil, err
		}
		var resultSquad model.Squad
		err := dao.FindOne("squad", name, &resultSquad)
		if err != nil {
			return nil, err
		}
		for j := range resultSquad.Devs {
			dev := resultSquad.Devs[j]
			update = bson.M{"$set": bson.M{"update_date": update_date, "active": b}}
			err = dao.Update("dev", dev, update)
			if err != nil {
				return nil, err
			}
		}
	}
	var findBU model.BU
	err = dao.FindOne("BU", name, &findBU)
	return &findBU, err
}
