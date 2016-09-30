package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
	"model"
	"encoding/json"
)

func TestInsertSquadToBUService(t *testing.T) {

}
func TestInsertSQuadToBUService_ShouldBeHaveSquadInsideBU_AfterInsert(t *testing.T) {
	assert := assert.New(t)
	unit := "dev"
	dev :=  model.Dev{Name:unit,Dev_id:"DevID",Update_date:100,Create_date:100,Active:"true"}
	target := "squad"
	squad := model.Squad{Name:target,Devs:[]model.Dev{},Update_date:100,Create_date:100,Active:"true"}
	findoneInsertDevtoSquadBefore = func(c string,name string,obj interface{})(error) {
		if c == "dev"{
			array := dev
			data,_ := json.Marshal(&array)
			json.Unmarshal(data,obj)
			return nil
		}else{
			array := squad
			data,_ := json.Marshal(&array)
			json.Unmarshal(data,obj)
			return nil
		}
	}
	updateInsertDevtoSquad = func(string,interface{},interface{})(error) {
		return nil
	}
	excepted := model.Squad{Name:target,Devs:[]model.Dev{dev},Update_date:100,Create_date:100,Active:"true"}
	findoneInsertDevtoSquadAfter = func(c string,name string,obj interface{})(error){
		array := excepted
		data,_ := json.Marshal(&array)
		json.Unmarshal(data,obj)
		return nil
	}
	actual,err := InsertDevToSquadService(unit,target)
	assert.Equal(&excepted,actual,"Insert Dev To Squad Servive Should return Squad that have dev inside")
	assert.Equal(nil,err,"Insert Dev To Squad Service Normal case Error should be nil")
}

func TestInsertDevToSquadService_FindSquadExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "dev"
	target := "squad"
	findoneInsertDevtoSquadBefore = func(c string,name string,obj interface{})(error) {
		if name == "squad" {
			return errors.New("Call DAO Fine one Return Error")
		}else{
			return nil
		}
	}
	actual,err := InsertDevToSquadService(unit,target)
	assert.Nil(actual,"When Dao Fine one Return Error result should be nil")
	assert.Error(err,"When Dao Fine one Return Error should be return error")


}

func TestInsertDevToSquadService_FindDevExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "dev"
	target := "squad"
	findoneInsertDevtoSquadBefore = func(c string,name string,obj interface{})(error){
		if name == "dev" {
			return errors.New("Call DAO Fine one Return Error")
		}else{
			return nil
		}
	}
	actual,err := InsertDevToSquadService(unit,target)
	assert.Nil(actual,"When Dao Fine one Return Error result should be nil")
	assert.Error(err,"When Dao Fine one Return Error should be return error")
}

func TestInsertDevToSquadService_UpdateExceptError_WhenMockDAO_ReturnError(t *testing.T) {
	assert := assert.New(t)
	unit := "dev"
	target := "squad"
	findoneInsertDevtoSquadBefore = func(string,string,interface{})(error){
		return nil
	}
	updateInsertDevtoSquad = func(string,interface{},interface{})(error) {
		return errors.New("Call DAO Update Return Error")
	}
	actual,err := InsertDevToSquadService(unit,target)
	assert.Nil(actual,"When Dao Update Return Error result should be nil")
	assert.Error(err,"When Dao Update Return Error should be return error")
}