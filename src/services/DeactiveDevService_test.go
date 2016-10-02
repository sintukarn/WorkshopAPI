package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"encoding/json"
	"errors"
)

func TestDeactiveDevService_ShouldBeUpdateNewBooleanAfterCallFunc(t *testing.T) {
	assert := assert.New(t)
	name := "bobo"
	excepted := model.Dev{Name:name,Dev_id:"DEVID",Create_date:100,Update_date:100,Active:"false"}
	findoneDeactiveDevBefore = func (c string,name string,obj interface{})(error) {
		result := model.Dev{Name:name,Dev_id:"DEVID",Create_date:100,Update_date:100,Active:"true"}
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	updateDeactiveDev = func (string ,interface{},interface{})error{
		return nil
	}
	findoneDeactiveDevAfter = func (c string,name string,obj interface{})(error) {
		result := model.Dev{Name:name,Dev_id:"DEVID",Create_date:100,Update_date:100,Active:"false"}
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	actual,err := DeactiveDevService(name,"false")
	assert.Equal(&excepted,actual,"DeactiveDevService Active Should change to false")
	assert.Nil(err,"DeactiveDevService in normal case error should return nil")
}

func TestDeactiveDevService_FindoneAfterError_ShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	findoneDeactiveDevAfter = func (string,string, interface{})(error){
		return errors.New("Find one Return some error")
	}
	actual,err := DeactiveDevService("test error","false")
	assert.Error(err,"DeativeDer if fine one error should return error")
	assert.Nil(actual,"Deactive find one error result should be nil")
}

func TestDeactiveDevService_UpdateError_ShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	updateDeactiveDev = func (string,interface{}, interface{})(error){
		return errors.New("Update Return some error")
	}
	actual,err := DeactiveDevService("test error","false")
	assert.Error(err,"DeativeDer if update error should return error")
	assert.Nil(actual,"Deactive update error result should be nil")
}

func TestDeactiveDevService_FindoneBeforeError_ShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	findoneDeactiveDevBefore = func (string,string, interface{})(error){
		return errors.New("Find one Return some error")
	}
	actual,err := DeactiveDevService("test error","false")
	assert.Error(err,"DeativeDer if fine one error should return error")
	assert.Nil(actual,"Deactive find one error result should be nil")
}