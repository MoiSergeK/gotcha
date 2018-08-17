package libs

import (

    "io/ioutil"

)

func View(view string) string {
	dat, err := ioutil.ReadFile("./app/views/" + view + ".html")
	if err != nil {
        panic(err)
    }
	return string(dat)
}

func JS(file string) string {
	dat, err := ioutil.ReadFile("./app/views/js/" + file)
	if err != nil {
        panic(err)
    }
	return string(dat)
}