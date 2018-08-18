package libs

import (
	"io/ioutil"
)

func ReadFileAsText(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	
	if err != nil { panic(err) }

	return string(file)
}