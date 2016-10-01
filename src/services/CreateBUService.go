package services

import (
	"dao"
	"model"
	"time"
)

var (
	create_date_func = time.Now().Unix
	update_date_func = time.Now().Unix
	insertCreateBU   = dao.Insert
	findoneCreateBU  = dao.FindOne
)

func CreateBUService(name string) (*model.BU, error) {
	create_date := create_date_func()
	update_date := update_date_func()
	bu := &model.BU{Name: name, Squads: []model.Squad{}, Create_date: create_date, Update_date: update_date, Active: "true"}
	err := insertCreateBU("BU", bu)
	if err != nil {
		return nil, err
	}
	var result model.BU
	err = findoneCreateBU("BU", name, &result)
	return &result, err
}
