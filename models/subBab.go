package models

type SubBab struct {
	Id				int			`gorm:"primaryKey"`
	Title 			string
	Order 			int
	ImageURL 		string
	InstructorName	string
	BabID 			int
  	Bab   			Bab
	Contents 		[]Content	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}