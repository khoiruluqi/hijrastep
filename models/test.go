package models

// for both quiz dan ujian

type Test struct {
	MinimumScore		int
	ContentID 			int		`gorm:"primaryKey;autoIncrement:false"`
	Content				Content `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Questions 			[]Question	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logs	 			[]ExamLog	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}