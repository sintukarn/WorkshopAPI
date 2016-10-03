package services

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"model"
	"testing"
)

func TestDeactiveSquadService_DeactiveSquadThatHaveDevInsideShouldBeFalseAll(t *testing.T) {
	assert := assert.New(t)
	name := "SQUAD"
	dev := model.Dev{Name: "DEV", Dev_id: "DEVID", Update_date: 100, Create_date: 100, Active: "true"}
	squad := model.Squad{Name: name, Devs: []model.Dev{dev}, Create_date: 100, Update_date: 100, Active: "true"}
	exceptedDev := model.Dev{Name: "DEV", Dev_id: "DEVID", Update_date: 100, Create_date: 100, Active: "false"}
	exceptedSquad := model.Squad{Name: name, Devs: []model.Dev{dev}, Create_date: 100, Update_date: 100, Active: "false"}
	findoneDeactiveSquadBefore = func(c string, name string, obj interface{}) error {
		result := squad
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	deacDev = func(string, string) (*model.Dev, error) {
		return &exceptedDev, nil
	}
	updateDeactiveSquad = func(string, interface{}, interface{}) error {
		return nil
	}
	findoneDeactiveSquadAfter = func(c string, name string, obj interface{}) error {
		result := exceptedSquad
		data, _ := json.Marshal(&result)
		json.Unmarshal(data, obj)
		return nil
	}
	actual, err := DeactiveSquadService(name, "false")
	assert.Equal(&exceptedSquad, actual, "Deactive Squad Should change Stauts Squad to false and Dev inside squad to false")
	assert.Nil(err, "Deactive Squad in normal case Should be return nil in error")
}

func TestDeactiveSquadService_UpdateErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	updateDeactiveSquad = func(string, interface{}, interface{}) error {
		return errors.New("Update in dao error")
	}
	actual, err := DeactiveSquadService("test error", "false")
	assert.Error(err, "Deactive Squad Should return Error when update error")
	assert.Nil(actual, "Deactive Squad Should Equal Nil When error")
}

func TestDeactiveSquadService_DeactiveDevErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	deacDev = func(string, string) (*model.Dev, error) {
		return nil, errors.New("Deactive dev in error")
	}
	actual, err := DeactiveSquadService("test error", "false")
	assert.Error(err, "Deactive Squad Should return Error when Deactive error")
	assert.Nil(actual, "Deactive Squad Should Equal Nil When error")
}

func TestDeactiveSquadService_FindOneErrorShouldReturnError(t *testing.T) {
	assert := assert.New(t)
	findoneDeactiveSquadBefore = func(string, string, interface{}) error {
		return errors.New("fine one dev in dao error")
	}
	actual, err := DeactiveSquadService("test error", "false")
	assert.Error(err, "fineone Squad Should return Error when Deactive error")
	assert.Nil(actual, "fineone Squad Should Equal Nil When error")
}
