package services

import (
	"dao"
	"model"
	"time"
)

var (
	insertCreateDev  = dao.Insert
	findoneCreateDev = dao.FindOne
)

func CreateDevService(id string, name string) (*model.Dev, error) {
	create_date := time.Now().Unix()
	update_date := time.Now().Unix()
	dev := &model.Dev{Dev_id: id, Name: name, Create_date: create_date, Update_date: update_date, Active: "true"}
	err := insertCreateDev("dev", dev)
	if err != nil {
		return nil, err
	}
	var result model.Dev
	err = findoneCreateDev("dev", name, &result)
	return &result, err
}
