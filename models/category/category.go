package category

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const categoriesUrl = "https://2ch.hk/makaba/mobile.fcgi?task=get_boards"

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetList() interface{} {
	resp, respErr := http.Get(categoriesUrl)

	if respErr != nil || resp.StatusCode > 205 {
		log.Fatal(respErr)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var categories map[string][]Category

	jsonErr := json.Unmarshal(byteValue, &categories)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return categories
}
