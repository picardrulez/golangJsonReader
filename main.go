package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type readStruct struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

func main() {
	res, err := http.Get("http://localhost:8000")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(res.Body)
	var jsonResponse readStruct
	err = decoder.Decode(&jsonResponse)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(jsonResponse.Version)
}
