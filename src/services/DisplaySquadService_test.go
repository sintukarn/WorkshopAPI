package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"encoding/json"
)

func TestDisplaySquadService_Normalcase_ShouldBeReturnArray2Squad_andHaveDevinside(t *testing.T) {
	assert := assert.New(t)
	name := "Universe"
	dev := model.Dev{Name:"Apolo",Dev_id:"DEVAPOLO",Update_date:100,Create_date:100,Active:"true"}
	squadM := model.Squad{Name:"M",Devs:[]model.Dev{dev},Update_date:100,Create_date:100,Active:"true"}
	squadN:= model.Squad{Name:"N",Devs:[]model.Dev{dev},Update_date:100,Create_date:100,Active:"false"}
	excepted := []model.BU{model.BU{Name:name,Squads:[]model.Squad{squadM,squadN},Update_date:100,Create_date:100,Active:"true"}}
	findallDisplaySquad = func(c string,obj interface{})(error) {
		array := excepted
		data,_ := json.Marshal(&array)
		json.Unmarshal(data,obj)
		return nil
	}
	actual,err := DisplaySquadService()
	assert.Equal(excepted,actual,"Display Squad Service Show be return model.BU when call method")
	assert.Nil(err,"Display Squad Shoud be return arror as nil in normal case")
}
