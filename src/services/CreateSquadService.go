package services

import (
	"model"
	"time"
	"dao"
)
var (
	insertCreateSquad = dao.Insert
	findoneCreateSquad = dao.FindOne
)

func CreateSquadService(name string)(*model.Squad,error)  {
	create_date := time.Now().Unix()
	update_date := time.Now().Unix()
	squad := &model.Squad{Name:name,Devs:[]model.Dev{},Create_date:create_date,Update_date:update_date,Active:"true"}
	err := insertCreateSquad("squad",squad)
	if err != nil {
		return nil,err
	}
	var result model.Squad
	err = findoneCreateSquad("squad",name,&result)
	return &result,err
}