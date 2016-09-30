package model

type BU struct {
	Name		string
	Squads		[]Squad
	Create_date	int64
	Update_date	int64
	Active		string
}

type Squad struct {
	Name		string
	Devs		[]Dev
	Create_date	int64
	Update_date	int64
	Active		string
}

type Dev struct {
	Dev_id		string
	Name		string
	Create_date	int64
	Update_date	int64
	Active		string
}