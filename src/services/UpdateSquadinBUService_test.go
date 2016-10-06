package services

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"model"
	"testing"
)

func TestUpdateSquadinBUService_ShouldBeReturnNil_HaveNameEqualSquadInsideBU(t *testing.T) {
	assert := assert.New(t)
	exceptedSquad := model.Squad{Name: "SQUAD", Devs: []model.Dev{}, Create_date: 100, Update_date: 100, Active: "true"}
	fineoneUpdateSquadinBUBefore = func(c string, name string, obj interface{}) error {
		squad := exceptedSquad
		data, _ := json.Marshal(&squad)
		json.Unmarshal(data, obj)
		return nil
	}
	fineAllUpdateSquadinBu = func(c string, obj interface{}) error {
		bu := []model.BU{model.BU{Name: "BU", Squads: []model.Squad{exceptedSquad}, Create_date: 100, Update_date: 100, Active: "true"}}
		data, _ := json.Marshal(&bu)
		json.Unmarshal(data, obj)
		return nil
	}
	updateUpdateSquadinBU = func(string, interface{}, interface{}) error {
		return nil
	}
	err := UpdateSquadinBUService("SQUAD")
	assert.Nil(err, "UpdateSquadinBUService Should be return error as nil")
}
