package models

type Bab struct {
	Id    int `gorm:"primaryKey"`
	Title string
	Order int
	SubBabs []SubBab `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
