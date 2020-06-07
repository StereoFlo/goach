package goach

import (
	"goach/models"
)

func GetCategoryList() []models.Category {
	return models.GetList()
}

func GetBoardById(boardId string) []models.BoardWithThreads {
	return models.GetBoardById(boardId)
}
