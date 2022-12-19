package repository

import (
	"app/db"
	"app/db/models"
	"context"
	"strconv"
)

func GetTask(ctx context.Context, taskId string) (res *models.Task, err error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		return
	}

	task, err := models.FindTask(ctx, db, taskIdInt)
	if err != nil {
		return nil, err
	}

	return task, err
}
