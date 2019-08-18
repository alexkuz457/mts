package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

//addTask создает новую задачу в БД с помощью CreateTask,
//возвращает пользователю GUID задачи
//передает управление в updateTask
func addTask(c echo.Context) (err error) {
	db, err := Database()

	if err != nil {
		return c.JSON(500, err)
	}

	t, err := CreateTask(db)

	if err != nil {
		return c.JSON(500, err)
	}

	r := new(PostResponse)
	r.GUID = t.GUID

	defer db.Close()
	defer FinishTask(t)
	defer UpdateTask(db, t)
	return c.JSON(http.StatusAccepted, r)
}

//CreateTask создает новую задачу в базе данных
func CreateTask(db *gorm.DB) (t *Task, err error) {
	t = new(Task)
	t.GUID = uuid.New()
	t.Status = "created"
	t.TimeStamp = time.Now()
	db.Create(t)
	return t, err
}

//UpdateTask обновляет статус задачи в БД после сообщения GUID пользователю
func UpdateTask(db *gorm.DB, t *Task) (*Task, error) {
	var err error
	t.TimeStamp = time.Now()
	t.Status = "running"

	db.Save(t)
	go FinishTask(t)
	return t, err
}

//FinishTask запускает новую горутину завершения задачи
// ничего не возвращает - возвращать некому
// безымянная функция используется для лучшего понимания кода материнской функции addTask
func FinishTask(t *Task) {
	go func(t *Task) {

		time.Sleep(time.Minute * 2)
		t.TimeStamp = time.Now()
		t.Status = "finished"
		db, _ := Database()
		db.Save(t)
		db.Close()
	}(t)

}
