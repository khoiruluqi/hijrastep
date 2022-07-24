package models

import "github.com/lib/pq"

type Question struct {
	Id       int `gorm:"primaryKey"`
	Order    int
	Question string
	Hint	 string
	Options  pq.StringArray `gorm:"type:varchar(256)[]"`
	Answer   int
	TestID   int
	Test     Test
}
