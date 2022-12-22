package repository

import (
	"app/db"
	"app/db/models"
	"context"
)

func GetTask(ctx context.Context, taskId int) (res *models.Task, err error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	task, err := models.FindTask(ctx, db, taskId)
	if err != nil {
		return nil, err
	}

	return task, err
}
