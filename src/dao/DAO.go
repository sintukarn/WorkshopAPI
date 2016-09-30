package dao

import (
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

var (
	url = "localhost:27017"
	dial = mgo.Dial
	marshall =  json.Marshal
)

func Insert(c string,input interface{})(error)  {
	session, err := dial(url)
	if err != nil {
		return errors.New("Cant connect database")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	cc := session.DB("API").C(c)
	err = cc.Insert(&input)
	if err != nil {
		return err
	}
	return nil
}

func Update(c string,find interface{},update interface{})(error) {
	session, err := dial(url)
	if err != nil {
		return errors.New("Cant connect database")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	cc := session.DB("API").C(c)
	err = cc.Update(&find,&update)
	if err != nil {
		return err
	}
	return nil
}

func FindOne(c string,name string,obj interface{})(error) {
	var result interface{}
	session, err := dial(url)
	if err != nil {
		return errors.New("Cant connect database")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	cc := session.DB("API").C(c)
	err = cc.Find(bson.M{"name":name}).One(&result)
	query,err := marshall(result)
	if err != nil {
		return errors.New(err.Error())
	}
	err = json.Unmarshal(query,&obj)
	return err
}

func FindAll(c string,obj interface{})(error){
	var result []interface{}
	session, err := dial(url)
	if err != nil {
		return errors.New("Cant connect database")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	cc := session.DB("API").C(c)
	err = cc.Find(nil).All(&result)
	query,err := marshall(&result)
	if err != nil {
		return errors.New(err.Error())
	}
	err = json.Unmarshal(query,obj)
	return err
}
