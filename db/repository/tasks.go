package repository

import (
	"app/db"
	"app/db/models"
	"log"
)

func GetTask(taskId string) (res *models.Task, err error) {
	db, err := db.Connect()
	if err != nil {
		log.Println("task:", "Connect")
		return nil, err
	}

	task := &models.Task{}
	rows, err := db.Query("SELECT * FROM tasks WHERE id = ?", taskId)
	if err != nil {
		log.Println("err", db)
		log.Println("err", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.EndDate, &task.Status, &task.Priority, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("task:", "err")
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println("task:", "row")
		return nil, err
	}

	return task, err
}
