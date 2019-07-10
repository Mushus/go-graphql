package main

import (
	"context"

	go_graphql "github.com/Mushus/go-graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type Resolver struct {
	db *gorm.DB
}

func (r *Resolver) Mutation() go_graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() go_graphql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateWebSite(ctx context.Context, input go_graphql.NewWebSite) (*go_graphql.WebSite, error) {
	entity := WebSite{
		Name: input.Name,
	}

	if err := r.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	return entity.ToRespModel(), nil
}
func (r *mutationResolver) CreateDocumument(ctx context.Context, input go_graphql.NewDocumument) (*go_graphql.Document, error) {
	var categories Categories
	r.db.Where("name in (?)", input.Categories).Find(&categories)

	entity := Document{
		Title:      input.Title,
		Body:       input.Body,
		Categories: categories,
	}
	if err := r.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return entity.ToRespModel(), nil
}
func (r *mutationResolver) CreateCategory(ctx context.Context, input go_graphql.NewCategory) (*go_graphql.Category, error) {
	entity := Category{
		Name: input.Name,
	}
	if err := r.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return entity.ToRespModel(), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) WebSite(ctx context.Context, id string) (*go_graphql.WebSite, error) {
	var entity WebSite
	db := r.db.Where("id = ?", id).First(&entity)
	if db.RecordNotFound() {
		return nil, nil
	}
	return entity.ToRespModel(), db.Error
}
func (r *queryResolver) Document(ctx context.Context, id string) (*go_graphql.Document, error) {
	var entity Document
	db := r.db.Where("id = ?", id).First(&entity)
	if db.RecordNotFound() {
		return nil, nil
	}
	return entity.ToRespModel(), db.Error
}
func (r *queryResolver) Category(ctx context.Context, id *string, name *string) (*go_graphql.Category, error) {
	if id == nil && name == nil {
		return nil, xerrors.New("at least you need to use one of id or name")
	}
	db := r.db

	if id != nil {
		db = db.Where("id = ?", *id)
	}

	if name != nil {
		db = db.Where("name = ?", *name)
	}

	var entity Category
	db = db.First(&entity)
	if db.RecordNotFound() {
		return nil, nil
	}

	return entity.ToRespModel(), db.Error
}
func (r *queryResolver) AllWebSites(ctx context.Context) ([]*go_graphql.WebSite, error) {
	panic("not implemented")
}
func (r *queryResolver) AllDocuments(ctx context.Context, input *go_graphql.QueryDocuments) ([]*go_graphql.Document, error) {
	panic("not implemented")
}
func (r *queryResolver) AllCategories(ctx context.Context, input *go_graphql.QueryCategories) ([]*go_graphql.Category, error) {
	panic("not implemented")
}
