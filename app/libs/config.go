package libs

import (
	"io/ioutil"
	"encoding/json"
)

type DBStruct struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Db string `json:"db"`
	Driver string `json:"driver"`
}

func DB() map[string]string {
	file, err := ioutil.ReadFile("./app/config/config.json")
	if err != nil { panic(err) }
	db := DBStruct{}
	json.Unmarshal(file, &db)
	return map[string]string{"user": db.User, "pass": db.Pass, "db": db.Db, "driver": db.Driver}
}