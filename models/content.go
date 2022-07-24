package models

type Content struct {
	Id					int		`gorm:"primaryKey"`
	Title 				string
	Order 				int
	DurationInMinute	int
	Type				string
	SubBabID 			int
	SubBab   			SubBab
	Logs	 			[]ContentLog	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}