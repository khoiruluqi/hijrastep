package models

import "time"

type ContentLog struct {
	ContentID     	int		`gorm:"primaryKey;autoIncrement:false"`
	Email     		string	`gorm:"primaryKey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	Content			Content
}