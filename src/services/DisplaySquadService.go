package services

import (
	"dao"
	"model"
)
var(
	findallDisplaySquad = dao.FindAll
)

func DisplaySquadService()([]model.BU,error){
	var result []model.BU
	err := findallDisplaySquad("BU",&result)
	return result,err
}