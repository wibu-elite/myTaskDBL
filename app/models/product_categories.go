package models

type ProductCategories struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ProductID    string `gorm:"size:36;index"`
	Product      Product
	CategoriesID string `gorm:"size:36;index"`
	Categories   Categories
}
