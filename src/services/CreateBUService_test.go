package services

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"model"
	"testing"
)

func TestCreateBUService(t *testing.T) {
	assert := assert.New(t)
	name := "Star"
	expected := model.BU{Name: name, Squads: []model.Squad{}, Create_date: 100, Update_date: 100, Active: "true"}
	update_date_func = func() int64 { return 100 }
	create_date_func = func() int64 { return 100 }
	insertCreateBU = func(c string, input interface{}) error {
		return nil
	}
	findoneCreateBU = func(c string, name string, obj interface{}) error {
		result := model.BU{Name: name, Squads: []model.Squad{}, Create_date: 100, Update_date: 100, Active: "true"}
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	actualBU, err := CreateBUService(name)
	assert.Equal(&expected, actualBU, "Create BU Object Should be create name same input")
	assert.Nil(err, "Create Bu Service normal case 'err' should be return nil")
}

func TestCreateBUService_WhenInsertCreateBuErr_ShouldBeReturnErr(t *testing.T) {
	assert := assert.New(t)
	insertCreateBU = func(string, interface{}) error {
		return errors.New("Insert Create BU error")
	}
	result, err := CreateBUService("error")
	if err != nil {
		assert.Error(err, "Should be return error when insert error")
	} else {
		assert.Nil(result, "When Create Bu Error result shuld be nil")
	}
}
