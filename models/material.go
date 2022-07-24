package models

type Material struct {
	VideoURL 			string
	AudioURL			string
	Text				string
	ContentID 			int		`gorm:"primaryKey;autoIncrement:false"`
	Content				Content `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}