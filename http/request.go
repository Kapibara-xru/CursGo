package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	Payload []interface{} `json:"payload"`
	Results []interface{} `json:"results"`
}
type Message struct {
	User_name string `json:"user_name"`
	Task      string `json:"task"`
	Results   result `json:"results"`
}

//Делаем GET запрос на сервер
func HttpGetRequest(url string) []interface{} {
	var mas []interface{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &mas)

	return mas
}

//Делаем POST запрос на сервер
func HttpPostRequst(nickName string, task string, payload []interface{}, results []interface{}) string {
	client := http.Client{}

	msg := &Message{
		User_name: nickName,
		Task:      task,
		Results: result{
			Payload: payload,
			Results: results,
		},
	}
	bytesRepresentation, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "http://116.203.203.76:3000/tasks/solution", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("Задача: %s \nStatus: %s \nResponse: %s", task, resp.Status, body)
}
