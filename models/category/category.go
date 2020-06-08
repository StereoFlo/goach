package category

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const categoriesUrl = "https://2ch.hk/makaba/mobile.fcgi?task=get_boards"

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetList() (map[string][]Category, error) {
	resp, respErr := http.Get(categoriesUrl)

	if respErr != nil || resp.StatusCode > 205 {
		if respErr == nil {
			return nil, errors.New("error code more than 205")
		}
		return nil, respErr
	}

	defer resp.Body.Close()

	jsonByte, _ := ioutil.ReadAll(resp.Body)
	var categories map[string][]Category

	jsonErr := json.Unmarshal(jsonByte, &categories)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return categories, nil
}
