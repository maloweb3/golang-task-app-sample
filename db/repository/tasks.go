package repository

import (
	"app/db"
	"app/db/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetTasks(ctx context.Context) (res []*models.Task, err error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}

	tasks, err := models.Tasks(qm.Select("id", "name", "description", "end_date", "status", "priority")).All(ctx, db)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

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
