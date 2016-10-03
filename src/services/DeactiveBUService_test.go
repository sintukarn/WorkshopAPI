package services

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"model"
	"testing"
)

func TestDeactiveBUService_ShouldBeReturnBUStatusisFalse(t *testing.T) {
	assert := assert.New(t)
	name := "BU"
	dev := model.Dev{Name: "DEV", Dev_id: "DEVID", Update_date: 100, Create_date: 100, Active: "true"}
	squad := model.Squad{Name: "SQUAD", Devs: []model.Dev{dev}, Create_date: 100, Update_date: 100, Active: "true"}
	bu := model.BU{Name: name, Squads: []model.Squad{squad}, Create_date: 100, Update_date: 100, Active: "true"}
	exceptedDev := model.Dev{Name: "DEV", Dev_id: "DEVID", Update_date: 100, Create_date: 100, Active: "false"}
	exceptedSquad := model.Squad{Name: "SQUAD", Devs: []model.Dev{exceptedDev}, Create_date: 100, Update_date: 100, Active: "false"}
	exceptedBU := model.BU{Name: name, Squads: []model.Squad{squad}, Create_date: 100, Update_date: 100, Active: "false"}
	fineoneDeactiveBUBefore = func(c string, name string, obj interface{}) error {
		result := bu
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	deacSquad = func(string, string) (*model.Squad, error) {
		return &exceptedSquad, nil
	}
	updateDeactiveBU = func(string, interface{}, interface{}) error {
		return nil
	}
	fineoneDeactiveBUAfter = func(c string, name string, obj interface{}) error {
		result := exceptedBU
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	actual, err := DeactiveBUService(name, "false")
	assert.Equal(&exceptedBU, actual, "Deactive BU Should change Stauts BU to false and Squad,dev inside squad to false")
	assert.Nil(err, "Deactive BU in normal case Should be return nil in error")
}

func TestDeactiveBUService_UpdateErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	updateDeactiveBU = func(string, interface{}, interface{}) error {
		return errors.New("Update in dao error")
	}
	actual, err := DeactiveBUService("test error", "false")
	assert.Error(err, "Deactive BU Should return Error when update error")
	assert.Nil(actual, "Deactive BU Should Equal Nil When error")
}

func TestDeactiveBUService_DeactiveDevErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	deacSquad = func(string, string) (*model.Squad, error) {
		return nil, errors.New("Deactive dev in error")
	}
	actual, err := DeactiveBUService("test error", "false")
	assert.Error(err, "Deactive BU Should return Error when Deactive error")
	assert.Nil(actual, "Deactive BU Should Equal Nil When error")
}

func TestDeactiveBUService_FindOneErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	fineoneDeactiveBUBefore = func(string, string, interface{}) error {
		return errors.New("fine one dev in dao error")
	}
	actual, err := DeactiveBUService("test error", "false")
	assert.Error(err, "fineone BU Should return Error when Deactive error")
	assert.Nil(actual, "fineone BU Should Equal Nil When error")
}
