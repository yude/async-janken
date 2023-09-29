package main

type Player struct {
	Uuid string `gorm:"primaryKey" json:"uuid"`
	Name string `json:"name"`
}

type Match struct {
	Id        string `gorm:"primaryKey" json:"id"`
	PlayerOne string `json:"player_1"`
	PlayerTwo string `json:"player_2"`
	Winner    string `json:"winner"`
}

type CurrentMatch struct {
	Id            string
	PlayerOneUuid string
	PlayerOneHand int
	PlayerTwoUuid string
	PlayerTwoHand int
	Done          bool
}

type JankenResult struct {
	Id     string `json:"id"`
	Winner string `json:"winner"`
	Done   bool   `json:"is_done"`
}

type JankenRequest struct {
	Uuid string `json:"uuid"`
	Hand int    `json:"hand"`
}

var IsFirstPlayer bool
var Current CurrentMatch

func InitMatch() {
	IsFirstPlayer = false
	Current = CurrentMatch{}
}
