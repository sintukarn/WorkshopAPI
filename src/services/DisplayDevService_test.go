package services

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"model"
	"testing"
)

func TestDisplayDevService_Normalcase_ShouldBeReturnArray2Dev(t *testing.T) {
	assert := assert.New(t)
	excepted := []model.Dev{model.Dev{Name: "A", Dev_id: "DEVA", Create_date: 100, Update_date: 100, Active: "true"}, model.Dev{Name: "B", Dev_id: "DEVB", Create_date: 100, Update_date: 100, Active: "true"}}
	findallDispplayDev = func(c string, obj interface{}) error {
		array := excepted
		data, _ := json.Marshal(&array)
		json.Unmarshal(data, obj)
		return nil
	}
	actual, err := DisplayDevService()
	assert.Equal(excepted, actual, "Display Dev Service Normal case excepted model.DEV should be eqaul result from Display Der service")
	assert.Nil(err, "Display Dev Service normal case Should be return error as nil")
}
