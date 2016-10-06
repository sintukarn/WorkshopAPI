package services

import (
	"dao"
	"gopkg.in/mgo.v2/bson"
	"model"
)

var (
	fineoneUpdateSquadinBUBefore = dao.FindOne
	fineAllUpdateSquadinBu       = dao.FindAll
	updateUpdateSquadinBU        = dao.Update
)

func UpdateSquadinBUService(name string) error {
	var resultSquad model.Squad
	err := fineoneUpdateSquadinBUBefore("squad", name, &resultSquad)
	if err != nil {
		return err
	}
	update_date := update_date_func()
	var resultBU []model.BU
	err = fineAllUpdateSquadinBu("BU", &resultBU)
	if err != nil {
		return err
	}
	for i := range resultBU {
		for j := range resultBU[i].Squads {
			if resultBU[i].Squads[j].Name == name {
				resultBU[i].Squads[j] = resultSquad
			}
			find := bson.M{"name": resultBU[i].Name}
			update := bson.M{"$set": bson.M{"squads": resultBU[i].Squads, "update_date": update_date}}
			err = updateUpdateSquadinBU("BU", find, update)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
