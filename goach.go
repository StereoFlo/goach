package goach

import (
	"goach/models/board"
	"goach/models/category"
)

func GetCategoryList() (map[string][]category.Category, error) {
	return category.GetList()
}

func GetBoardById(boardId string) (*board.ThreadsList, error) {
	return board.GetBoardById(boardId)
}
