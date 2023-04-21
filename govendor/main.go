package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

func main() {
	go interval()
	// eksekusi selama 300 detik
	time.Sleep(time.Second * 300)
}

func interval() {
	// eksekusi setiap 15 detik
	for range time.Tick(time.Second * 15) {
		waterWind()
	}
}

func waterWind() {
	water := gubrak.RandomInt(1, 100)
	wind := gubrak.RandomInt(1, 100)

	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/posts",
		bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	conditionStatus(water, wind)
	fmt.Println("")
}

func conditionStatus(water, wind int) {

	if water <= 5 {
		log.Println("status water : aman")
	} else if water >= 6 && water <= 8 {
		log.Println("status water : siaga")
	} else {
		log.Println("status water : bahaya")
	}

	if wind <= 6 {
		log.Println("status wind : aman")
	} else if wind >= 7 && water <= 15 {
		log.Println("status wind : siaga")
	} else {
		log.Println("status wind : bahaya")
	}
}
