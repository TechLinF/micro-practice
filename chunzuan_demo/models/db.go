package models

import (
	"time"
)

type Task struct {
	Id          string    `xorm:"not null pk unique VARCHAR(45)"`
	OwnerId     string    `xorm:"not null unique VARCHAR(45)"`
	Title       string    `xorm:"VARCHAR(45)"`
	Description string    `xorm:"VARCHAR(45)"`
	CreateTime  time.Time `xorm:"DATETIME created"`
	UpdateTime  time.Time `xorm:"DATETIME updated"`
	Version     int       `xorm:"version"`
}
