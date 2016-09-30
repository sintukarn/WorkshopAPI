package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"errors"
	"encoding/json"
)

func TestCreateDevService(t *testing.T) {
	assert := assert.New(t)
	name := "Somchai"
	id := "DEV12345"
	expected := model.Dev{Name:name,Dev_id:id,Create_date:100,Update_date:100,Active:"true"}
	update_date_func = func()(int64) {return 100}
	create_date_func = func()(int64) {return 100}
	insertCreateDev = func(c string,input interface{})(error){
		return nil
	}
	findoneCreateDev = func(c string,name string,obj interface{})(error){
		result := &model.Dev{Name:name,Dev_id:"DEV12345",Create_date:100,Update_date:100,Active:"true"}
		data,_ := json.Marshal(&result)
		json.Unmarshal(data,obj)
		return nil
	}
	actual,err := CreateDevService(id,name)
	assert.Equal(&expected,actual,"Create Dev Object Should be create name same input")
	assert.Nil(err,"Create Dev Service normal case 'err' should be return nil")
}

func TestCreateDevService_DAO_Insert_ReturnError(t *testing.T) {
	assert := assert.New(t)
	name := "Somchai"
	id := "DEV12345"
	insertCreateDev = func(c string, input interface{}) (error) {
		return errors.New("Error nil pointer")
	}
	actual, err := CreateDevService(id, name)
	if err != nil {
		assert.Error(err, "Create Dev return error when insert return error")
	}else {
		assert.Fail(actual.Name,"Create Dev Service Should be error when insert error")
	}
}
