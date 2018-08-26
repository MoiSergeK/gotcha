package libs

import (
	// "encoding/json"
	"net/http"
	"fmt"
	"bytes"
)

func PushNotification(receiverId int, message string) {
	APIKey := "CCg6GBdvF0t3gILMhnqeBiBPV3zOi4truVN4ZcXZ6Ca4Ws9F2q4EME3h0OEl1IXf9yWe9SYztVhmytgTXsQE"
	appId := "asdsad"
	
	// jsonString, err := json.Marshal(
	// 	map[string]map[string]string{
	// 		"request": map[string]string{
	// 			"application": appId,
	// 			"auth": APIKey,
	// 			"notifications": [
	// 				map[string]string{
	// 					"send_date": "now",
	// 					"content": message }]}}
	// )

	jsonString := "{\"request\":{\"application\":\""+appId+"\",\"auth\":\""+APIKey+"\",\"notification\":[{\"send_date\":\"now\",\"content\":\""+message+"\"}]}}"

	// if err != nil { panic(err) }

	url := "https://cp.pushwoosh.com/json/1.3/createMessage"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonString)))
    req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}

    resp, err := client.Do(req)
	
	if err != nil { panic(err) }
	
	defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
}