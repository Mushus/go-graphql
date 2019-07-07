package main

import (
	"github.com/jinzhu/gorm"
)

type WebSite struct {
	gorm.Model
	Name      string
	Documents []Document
}

type Document struct {
	gorm.Model
	Name     string
	Body     string
	Category []Category `gorm:"many2many:document_categories;"`
}

type Category struct {
	gorm.Model
	Name      string
	Documents []Document `gorm:"many2many:document_categories;"`
}
