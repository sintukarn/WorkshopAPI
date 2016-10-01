package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"encoding/json"
	"errors"
)

func TestInsertSquadToBUService(t *testing.T) {

}
func TestInsertSquadToBUService_ShouldBeHaveSquadInsideBU_AfterInsert(t *testing.T) {
	assert := assert.New(t)
	unit := "squad"
	squad :=  model.Squad{Name:unit,Devs:[]model.Dev{},Update_date:100,Create_date:100,Active:"true"}
	target := "BU"
	bu := model.BU{Name:target,Squads:[]model.Squad{},Update_date:100,Create_date:100,Active:"true"}
	findoneInsertSquadtoBUBefore = func(c string,name string,obj interface{})(error) {
		if c == "squad"{
			array := squad
			data,_ := json.Marshal(&array)
			json.Unmarshal(data,obj)
			return nil
		}else{
			array := bu
			data,_ := json.Marshal(&array)
			json.Unmarshal(data,obj)
			return nil
		}
	}
	updateInsertSquadtoBU = func(string,interface{},interface{})(error) {
		return nil
	}
	excepted := model.BU{Name:target,Squads:[]model.Squad{squad},Update_date:100,Create_date:100,Active:"true"}
	findoneInsertSquadtoBUAfter = func(c string,name string,obj interface{})(error){
		array := excepted
		data,_ := json.Marshal(&array)
		json.Unmarshal(data,obj)
		return nil
	}
	actual,err := InsertSquadToBUService(unit,target)
	assert.Equal(&excepted,actual,"Insert Squad To BU Servive Should return BU that have Squad inside")
	assert.Equal(nil,err,"Insert Squad To BU Service Normal case Error should be nil")
}

func TestInsertSquadToBUService_FindBUExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "squad"
	target := "bu"
	findoneInsertSquadtoBUBefore = func(c string,name string,obj interface{})(error) {
		if name == "bu" {
			return errors.New("Call DAO Fine one Return Error")
		}else{
			return nil
		}
	}
	actual,err := InsertSquadToBUService(unit,target)
	assert.Nil(actual,"When Dao Fine one Return Error result should be nil")
	assert.Error(err,"When Dao Fine one Return Error should be return error")
}

func TestInsertSquadToBUService_FindSquadExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "squad"
	target := "bu"
	findoneInsertSquadtoBUBefore = func(c string,name string,obj interface{})(error){
		if name == "squad" {
			return errors.New("Call DAO Fine one Return Error")
		}else{
			return nil
		}
	}
	actual,err := InsertSquadToBUService(unit,target)
	assert.Nil(actual,"When Dao Fine one Return Error result should be nil")
	assert.Error(err,"When Dao Fine one Return Error should be return error")
}
func TestInsertSquadToBUBUService_UpdateExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "squad"
	target := "bu"
	findoneInsertSquadtoBUBefore = func(string,string,interface{})(error){
		return nil
	}
	updateInsertSquadtoBU = func(string,interface{},interface{})(error) {
		return errors.New("Call DAO Update Return Error")
	}
	actual,err := InsertSquadToBUService(unit,target)
	assert.Nil(actual,"When Dao Update Return Error result should be nil")
	assert.Error(err,"When Dao Update Return Error should be return error")
}


