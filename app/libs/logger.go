package libs

import (
	"os"
	"fmt"
	"time"
)

func Log(data string) {
	file, err := os.OpenFile("./app/log/log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil{
		fmt.Println(err)
		os.Exit(1) 
	}

	t := time.Now().Format(time.RFC3339)

	file.WriteString(t + ": " + data + "\n");

	defer file.Close() 
}

func Println(message string) {
	fmt.Println(">>>" + message)
}