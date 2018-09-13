package libs

import (
	"io/ioutil"
	"encoding/json"
)

var cfg map[string]interface{} = make(map[string]interface{})

func LoadConfig() {
	file, err := ioutil.ReadFile("./app/config/config.json")
	
	if err != nil { panic(err) }
	json.Unmarshal(file, &cfg)
}

func CfgDB() map[string]string {
	if len(cfg) == 0 { LoadConfig() }

	db := make(map[string]string)

	for key, value := range cfg {
		if(key == "db"){
			for k,v := range value.(map[string]interface{}) {
				db[k] = v.(string)
			}
		}
	}
	
	return map[string]string{"user": db["user"], "pass": db["pass"], "db": db["db"], "driver": db["driver"]}
}

func CfgSecret() string {
	if len(cfg) == 0 { LoadConfig() }

	for key, value := range cfg {
		if(key == "secret"){
			return value.(string)
		}
	}

	return ""
}