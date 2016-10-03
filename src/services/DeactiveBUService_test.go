package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"encoding/json"
)

func TestDeactiveBUService_ShouldBeReturnBUStatusisFalse(t *testing.T) {
	assert := assert.New(t)
	name := "BU"
	dev := model.Dev{Name:"DEV",Dev_id:"DEVID",Update_date:100,Create_date:100,Active:"true"}
	squad := model.Squad{Name:"SQUAD",Devs:[]model.Dev{dev},Create_date:100,Update_date:100,Active:"true"}
	bu := model.BU{Name:name,Squads:[]model.Squad{squad},Create_date:100,Update_date:100,Active:"true"}
	exceptedDev := model.Dev{Name:"DEV",Dev_id:"DEVID",Update_date:100,Create_date:100,Active:"false"}
	exceptedSquad := model.Squad{Name:"SQUAD",Devs:[]model.Dev{exceptedDev},Create_date:100,Update_date:100,Active:"false"}
	exceptedBU := model.BU{Name:name,Squads:[]model.Squad{squad},Create_date:100,Update_date:100,Active:"false"}
	fineoneDeactiveBUBefore = func(c string,name string,obj interface{})(error){
		result := bu
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	deacSquad = func(string,string)(*model.Squad,error){
		return &exceptedSquad,nil
	}
	updateDeactiveBU = func(string,interface{},interface{})(error){
		return nil
	}
	fineoneDeactiveBUAfter = func(c string,name string,obj interface{})(error){
		result := exceptedBU
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	actual,err := DeactiveBUService(name,"false")
	assert.Equal(&exceptedBU,actual,"Deactive BU Should change Stauts BU to false and Squad,dev inside squad to false")
	assert.Nil(err,"Deactive BU in normal case Should be return nil in error")
}
