package model

import "time"

type Todo struct {
	Id         int        `gorm:"primary_key" json:"id"`
	TaskName   string     `gorm:"column:task_name" json:"taskName"`
	TaskStatus string     `gorm:"column:task_status" json:"taskStatus"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
