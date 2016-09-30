package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"errors"
	"encoding/json"
)

func TestInsertBU_ShouldBe_Equal_Find_ModelBU(t *testing.T) {
	assert := assert.New(t)
	c := "BU"
	expected := model.BU{Name:"Star",Squads:[]model.Squad{},Create_date:100,Update_date:100,Active:"true" }
	_ = Insert(c,expected)
	var actual model.BU
	_ = FindOne(c,"Star",&actual)
	assert.Equal(expected,actual,"Insert BU(Star) into mongoDB and find this name as inserted")
}

func TestInsertSquad_ShouldBe_Equal_Find_ModelSquad(t *testing.T) {
	assert := assert.New(t)
	c := "squad"
	expected := model.Squad{Name:"Animal",Devs:[]model.Dev{},Create_date:100,Update_date:100,Active:"true" }
	_ = Insert(c,expected)
	var actual model.Squad
	_ = FindOne(c,"Animal",&actual)
	assert.Equal(expected,actual,"Insert Squad(Animal) into mongoDB and find this name as inserted")
}

func TestInsertDev_ShouldBe_Equal_Find_ModelDev(t *testing.T) {
	assert := assert.New(t)
	c := "dev"
	expected := model.Dev{Name:"Somchai",Dev_id:"DEV12345",Create_date:100,Update_date:100,Active:"true" }
	_ = Insert(c,expected)
	var actual model.Dev
	_ = FindOne(c,"Somchai",&actual)
	assert.Equal(expected,actual,"Insert Dev(People) into mongoDB and find this name as inserted")
}

func TestInsert_Input_NilPointerInterface(t *testing.T){
	assert := assert.New(t)
	c := "BU"
	expected := Insert(c,nil)
	assert.Error(expected,"Show Error when put nill pointer")
}

func TestFindAll_InDevTable_ShouldBe_ReturnDevArray(t *testing.T) {
	assert := assert.New(t)
	c := "dev"
	expected := []model.Dev{model.Dev{Name:"Somchai",Dev_id:"DEV12345",Create_date:100,Update_date:100,Active:"true"},model.Dev{Name:"Somsri",Dev_id:"DEV56789",Create_date:100,Update_date:100,Active:"true"}}
	moreInsert := model.Dev{Name:"Somsri",Dev_id:"DEV56789",Create_date:100,Update_date:100,Active:"true" }
	_ = Insert(c,moreInsert)
	var actual []model.Dev
	_ = FindAll(c,&actual)
	assert.Equal(expected,actual,"FindAll Dev Should Be Eqaul 2 Dev")
}

func TestUpdateBU_ChangeUpdate_Active(t *testing.T) {
	assert := assert.New(t)
	expected := model.BU{Name:"Star",Squads:[]model.Squad{},Create_date:100,Update_date:100,Active:"false" }
	c := "BU"
	find := bson.M{"name":"Star"}
	update := bson.M{"$set":bson.M{"active":"false"}}
	_ = Update(c,find,update)
	var actual model.BU
	_ = FindOne(c,"Star",&actual)
	assert.Equal(expected,actual,"Update BU(Star) and active value in db change to false")
}

func TestUpdateSquad_ChangeUpdate_Active(t *testing.T) {
	assert := assert.New(t)
	expected := model.Squad{Name:"Animal",Devs:[]model.Dev{},Create_date:100,Update_date:100,Active:"false" }
	c := "squad"
	find := bson.M{"name":"Animal"}
	update := bson.M{"$set":bson.M{"active":"false"}}
	_ = Update(c,find,update)
	var actual model.Squad
	_ = FindOne(c,"Animal",&actual)
	assert.Equal(expected,actual,"Update Squad(Animal) and active value in db change to false")
}

func TestUpdateDev_WithNilPointer(t *testing.T){
	assert := assert.New(t)
	c := "dev"
	find := bson.M{"name":"PaPaya"}
	expected := Update(c,find,nil)
	assert.Error(expected,"Update with nil pointer should be return error")
}

func TestUpdateDev_ChangeUpdate_Active(t *testing.T) {
	assert := assert.New(t)
	expected := model.Dev{Name:"Somchai",Dev_id:"DEV12345",Create_date:100,Update_date:100,Active:"false" }
	c := "dev"
	find := bson.M{"name":"Somchai"}
	update := bson.M{"$set":bson.M{"active":"false"}}
	_ = Update(c,find,update)
	var actual model.Dev
	_ = FindOne(c,"Somchai",&actual)
	assert.Equal(expected,actual,"Update Dev(Somchai) and active value in db change to false")
}

func TestClearDBAfterTest(t *testing.T){
	assert := assert.New(t)
	session, _ := mgo.Dial("localhost:27017")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	cc := session.DB("API").C("BU")
	selector := bson.M{"name":"Star"}
	err := cc.Remove(selector)
	assert.Equal(err,nil,"Clear BU Data Finish")
	cc = session.DB("API").C("squad")
	selector = bson.M{"name":"Animal"}
	err = cc.Remove(selector)
	assert.Equal(err,nil,"Clear Squad Data Finish")
	cc = session.DB("API").C("dev")
	selector = bson.M{"name":"Somchai"}
	err = cc.Remove(selector)
	cc = session.DB("API").C("dev")
	selector = bson.M{"name":"Somsri"}
	err = cc.Remove(selector)
	assert.Equal(nil,err,"Clear Dev Data Finish")
}


func TestJsonMarshallError(t *testing.T){
	assert := assert.New(t)
	c := "dev"
	marshall = func(interface{})([]byte,error){
		return nil,errors.New("Marshall Error")
	}
	actual := FindOne(c,"name",nil)
	assert.Error(actual,"When mock Marshall should be except connect error")
	actual = FindAll(c,nil)
	assert.Error(actual,"When mock Marshall should be except connect error")
	marshall = func(v interface{})([]byte,error){
		return json.Marshal(v)
	}
}

func TestLostConnectionWhenConnect(t *testing.T){
	assert := assert.New(t)
	c := "dev"
	dial = func(string)(*mgo.Session,error){
		return nil,errors.New("Lost Connection")
	}
	actual := Insert(c,"name")
	assert.Error(actual,"When mock connect should be except connect error")
	actual = Update(c,nil,nil)
	assert.Error(actual,"When mock connect should be except connect error")
	actual = FindOne(c,"name",nil)
	assert.Error(actual,"When mock connect should be except connect error")
	actual = FindAll(c,nil)
	assert.Error(actual,"When mock connect should be except connect error")
}

