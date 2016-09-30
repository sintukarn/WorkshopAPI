package services

import (
	"model"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"errors"
)

func TestCreateSquadService(t *testing.T) {
	assert := assert.New(t)
	name := "Animal"
	expected := model.Squad{Name:name,Devs:[]model.Dev{},Create_date:100,Update_date:100,Active:"true"}
	update_date_func = func()(int64) {return 100}
	create_date_func = func()(int64) {return 100}
	insertCreateSquad = func(c string,input interface{})(error){
		return nil
	}
	findoneCreateSquad = func(c string,name string,obj interface{})(error){
		result := &model.Squad{Name:name,Devs:[]model.Dev{},Create_date:100,Update_date:100,Active:"true"}
		data,_ := json.Marshal(&result)
		json.Unmarshal(data,obj)
		return nil
	}
	actualSquad,err :=CreateSquadService(name)
	assert.Equal(&expected,actualSquad,"Create Dev Object Should be Equad Expected")
	assert.Nil(err,"Create Squad Service normal case 'err' should be return nil")
}

func TestCreateSquadService_WhenInsertCreateSquadErr_ShouldBeReturnErr(t *testing.T) {
	assert := assert.New(t)
	insertCreateSquad = func(string,interface{})(error) {
		return errors.New("Insert Create Squad error")
	}
	result,err := CreateSquadService("error")
	if err != nil {
		assert.Error(err, "Should be return error when insert error")
	}else{
		assert.Equal(nil,result,"When Create Squad Error result shuld be nil")
	}
}

