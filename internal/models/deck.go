package models

type Deck struct {
	ID       string `gorm:"primaryKey"`
	Shuffled bool
	Cards    Cards `gorm:"type:json"`
}
