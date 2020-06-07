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
	Comment    string      `json:"comment"`
	Lasthit    int32       `json:"lasthit"`
	Num        json.Number `json:"num"`
	PostsCount int32       `json:"posts_count"`
	Score      float32     `json:"score"`
	Views      int32       `json:"views"`
	Timestamp  int32       `json:"timestamp"`
	Subject    string      `json:"subject"`
}

type BoardWithThreads struct {
	Id           string   `json:"board"`
	Threads      []Thread `json:"threads"`
	ThreadsCount int      `json:"threads_count"`
}

func GetBoardById(boardId string) BoardWithThreads {
	resp, respErr := http.Get(fmt.Sprintf(threadsUrl, boardId))

	if respErr != nil || resp.StatusCode > 205 {
		log.Fatal(respErr)
	}

	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)
	var BoardWithThreads BoardWithThreads

	jsonErr := json.Unmarshal(byteValue, &BoardWithThreads)
	BoardWithThreads.ThreadsCount = len(BoardWithThreads.Threads)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return BoardWithThreads
}
