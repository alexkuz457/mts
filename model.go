package main

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Task структура задачи
type Task struct {
	GUID      uuid.UUID `gorm:"primary_key"` //GUID уникальный идентификатор
	Status    string    //Status статус
	TimeStamp time.Time //TimeStamp - время создания/последнего изменения
}

//GetResponse структура ответа запросу get
type GetResponse struct {
	Status    string    `json:"status"`    //Status статус
	TimeStamp time.Time `json:"timestamp"` //TimeStamp - время создания/время последнего изменения
}

//PostResponse структура ответа запросу post
type PostResponse struct {
	GUID uuid.UUID `json:"guid"`
}
