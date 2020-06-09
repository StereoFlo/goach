package board

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

type ThreadsList struct {
	Id           string   `json:"board"`
	Threads      []Thread `json:"threads"`
	ThreadsCount int      `json:"threads_count"`
}

func GetBoardById(boardId string) (*ThreadsList, error) {
	resp, respErr := http.Get(fmt.Sprintf(threadsUrl, boardId))

	if respErr != nil || resp.StatusCode > 205 {
		if respErr == nil {
			return nil, errors.New("status code" + string(resp.StatusCode))
		}
		return nil, respErr
	}

	defer resp.Body.Close()

	jsonValue, _ := ioutil.ReadAll(resp.Body)
	var ThreadsList ThreadsList

	jsonErr := json.Unmarshal(jsonValue, &ThreadsList)
	ThreadsList.ThreadsCount = len(ThreadsList.Threads)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return &ThreadsList, nil
}
