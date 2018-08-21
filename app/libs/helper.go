package libs

import (
	"io/ioutil"
)

func ReadFileAsText(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	
	if err != nil { panic(err) }

	return string(file)
}

func Map2JSONString(o map[string]string) string {
	json := "{"

	for key, value := range o {
		json += "\"" + string(key) + "\":" + "\"" + string(value) + "\","
	}

	return json + "}"
}