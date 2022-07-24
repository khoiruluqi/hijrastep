package models

import (
	"time"

	"github.com/lib/pq"
)

type ExamLog struct {
	TestID     		 int		`gorm:"primaryKey;autoIncrement:false"`
	Email     		 string		`gorm:"primaryKey"`
	Score            int
	CorrectAnswerNum int
	Answers          pq.Int32Array `gorm:"type:integer[]"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Test             Test
}
