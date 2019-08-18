package main

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

//getTask возвращаяет статус задачи и таймстамп по GUID
//если не находит задачу возвращаяет 404
//если переданы неверрные параметры возвращает 400
func getTask(c echo.Context) (err error) {
	db, err := Database()

	//не удалось подключиться к бд
	if err != nil {
		return c.JSON(500, err)
	}
	r := new(GetResponse)
	id, err := uuid.Parse(c.Param("id"))

	//не удалось распарсить uuid. Возвращает 400, если передан не GUID
	if err != nil {
		return c.JSON(400, r)
	}
	t := new(Task)
	t.GUID = id
	db.First(t)

	//Задача по переданному guid не существует 	Возвращает 404
	if t.Status == "" {
		return c.JSON(404, r)
	}

	db.Close()

	r.Status = t.Status
	r.TimeStamp = t.TimeStamp

	return c.JSON(200, r)
}
