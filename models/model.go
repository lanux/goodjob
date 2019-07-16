package models

import (
	"time"
)

type Model struct {
	Int64Id
	WithUpdate
	WithCreate
}

type IntId struct {
	Id int `gorm:"primary_key"`
}

type Int64Id struct {
	Id int64 `gorm:"primary_key"`
}

type WithCreate struct {
	CreateTime time.Time
	CreateUser int64
}

type WithUpdate struct {
	UpdateTime time.Time
	UpdateUser int64
}
