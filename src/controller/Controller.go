package controller

import (
	"model"
	"github.com/gin-gonic/gin"
	"services"
	"net/http"
)

func CreateBU(c *gin.Context)()  {
	var json model.Name
	if c.BindJSON(&json) == nil {
		result,err := services.CreateBUService(json.Name)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func CreateSquad(c *gin.Context)() {
	var json model.Name
	if c.BindJSON(&json) == nil {
		result,err := services.CreateSquadService(json.Name)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func CreateDev(c *gin.Context)() {
	var json model.Name
	if c.BindJSON(&json) == nil {
		id := services.GenerateID("DEV")
		result,err := services.CreateDevService(id,json.Name)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func InsertSquadToBU(c *gin.Context) {
	var json model.Insert
	if c.BindJSON(&json) == nil {
		result,err := services.InsertSquadToBUService(json.Unit,json.Target)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func InsertDevToSquad(c *gin.Context) {
	var json model.Insert
	if c.BindJSON(&json) == nil {
		result,err := services.InsertDevToSquadService(json.Unit,json.Target)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func DeactiveBU(c *gin.Context){
	var json model.Status
	if c.BindJSON(&json) == nil {
		result,err := services.DeactiveBUService(json.Name,json.Status)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func DeactiveSquad(c *gin.Context){
	var json model.Status
	if c.BindJSON(&json) == nil {
		result,err := services.DeactiveSquadService(json.Name,json.Status)
		if(err==nil){
			c.JSON(http.StatusOK, gin.H{"Result": result})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
		}
	}
}

func DisplayDev(c *gin.Context){
	result,err := services.DisplayDevService()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"Result": result})
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
	}
}

func DisplaySquad(c *gin.Context){
	result,err := services.DisplaySquadService()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"Result": result})
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{"Error":  err.Error()})
	}
}