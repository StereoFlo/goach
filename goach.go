package goach

import (
	"goach/models/board"
	"goach/models/category"
)

func GetCategoryList() interface{} {
	return category.GetList()
}

func GetBoardById(boardId string) board.BoardWithThreads {
	return board.GetBoardById(boardId)
}
