package main

import (
	"strconv"

	go_graphql "github.com/Mushus/go-graphql"
	"github.com/jinzhu/gorm"
)

type (
	WebSite struct {
		gorm.Model
		Name      string
		Documents Documents
	}

	Document struct {
		gorm.Model
		WebSiteID  uint
		Title      string
		Body       string
		Categories Categories `gorm:"many2many:document_categories;"`
	}

	Category struct {
		gorm.Model
		Name      string
		Documents Documents `gorm:"many2many:document_categories;"`
	}

	Documents  []Document
	Categories []Category
)

func (w *WebSite) ToRespModel() *go_graphql.WebSite {
	if w == nil {
		return nil
	}
	return &go_graphql.WebSite{
		ID:        strconv.FormatUint(uint64(w.Model.ID), 10),
		Name:      w.Name,
		Documents: w.Documents.ToRespModel(),
	}
}

func (d Documents) ToRespModel() []*go_graphql.Document {
	resp := make([]*go_graphql.Document, 0, len(d))
	for _, document := range d {
		resp = append(resp, document.ToRespModel())
	}
	return resp
}

func (d *Document) ToRespModel() *go_graphql.Document {
	if d == nil {
		return nil
	}
	return &go_graphql.Document{
		ID:         strconv.FormatUint(uint64(d.Model.ID), 10),
		Title:      d.Title,
		Body:       d.Body,
		Categories: d.Categories.ToRespModel(),
	}
}

func (c Categories) ToRespModel() []*go_graphql.Category {
	resp := make([]*go_graphql.Category, 0, len(c))
	for _, category := range c {
		resp = append(resp, category.ToRespModel())
	}
	return resp
}

func (c *Category) ToRespModel() *go_graphql.Category {
	if c == nil {
		return nil
	}
	return &go_graphql.Category{
		ID:        strconv.FormatUint(uint64(c.Model.ID), 10),
		Name:      c.Name,
		Documents: c.Documents.ToRespModel(),
	}
}
