package models

import "time"

type ProductImages struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ProductID  string `gorm:"size:36;index"`
	Product    Product
	Path       string `gorm:"type:text"`
	ExtraLarge string `gorm:"size:255"`
	Large      string `gorm:"size:255"`
	Medium     string `gorm:"size:255"`
	Small      string `gorm:"size:255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
