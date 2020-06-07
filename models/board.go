package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const threadsUrl = "https://2ch.hk/%s/threads.json"

type Thread struct {
	Comment    string `json:"comment"`
	Lasthit    string `json:"lasthit"`
	Num        string `json:"num"`
	PostsCount string `json:"posts_count"`
	Score      string `json:"score"`
	Views      string `json:"views"`
	Timestamp  string `json:"timestamp"`
	Subject    string `json:"subject"`
}

type BoardWithThreads struct {
	Id      string   `json:"board"`
	Threads []Thread `json:"threads"`
}

func GetBoardById(boardId string) []BoardWithThreads {
	resp, respErr := http.Get(fmt.Sprintf(threadsUrl, boardId))

	if respErr != nil {
		log.Fatal(respErr)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var BoardWithThreads []BoardWithThreads

	jsonErr := json.Unmarshal(byteValue, &BoardWithThreads)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return BoardWithThreads
}
