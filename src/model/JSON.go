package model

type Name struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type Insert struct {
	Unit   string `form:"unit" json:"unit" binding:"required"`
	Target string `form:"target" json:"target" binding:"required"`
}

type Status struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
}
