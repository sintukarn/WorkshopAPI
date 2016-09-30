package services

import (
	"dao"
	"model"
)
var(
	findallDispplayDev = dao.FindAll
)

func DisplayDevService()([]model.Dev,error){
	var result []model.Dev
	err := findallDispplayDev("dev",&result)
	return result,err
}
