package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const categoriesUrl = "https://2ch.hk/makaba/mobile.fcgi?task=get_boards"

type Board struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Name       string  `json:"name"`
	BoardCount int32   `json:"board_count"`
	Boards     []Board `json:"boards"`
}

func GetList() []Category {
	resp, respErr := http.Get(categoriesUrl)

	if respErr != nil {
		log.Fatal(respErr)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var categories []Category

	jsonErr := json.Unmarshal(byteValue, &categories)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return categories
}