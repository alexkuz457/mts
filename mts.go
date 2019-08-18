package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

//Database инициализирует и возвращает соединение с postgres
func Database() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=mts password=qwerty12 sslmode=disable")
	return db, err
}

func main() {
	e := echo.New()
	e.POST("/task", addTask)
	e.GET("/task/:id", getTask)
	e.Logger.Fatal(e.Start(":9000"))
}
