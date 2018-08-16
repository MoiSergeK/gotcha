package libs

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func GET(url string) map[string]interface{} {
	// if params.size > 0 {
	// 	s := '?'
	// 	for key, value := range params { s += key + '=' + value + '&' }
	// }

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var info map[string]interface{}
    json.Unmarshal([]byte(string(body)),&info)

	return info
}

// func POST(url string, params [string]) {

// }