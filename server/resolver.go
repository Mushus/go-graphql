package main

import (
	"context"

	go_graphql "github.com/Mushus/go-graphql"
	"github.com/jinzhu/gorm"
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
	return nil, nil
}
func (r *mutationResolver) CreateDocumument(ctx context.Context, input go_graphql.NewDocumument) (*go_graphql.Document, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCategory(ctx context.Context, input go_graphql.NewCategory) (*go_graphql.Category, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) WebSites(ctx context.Context) ([]*go_graphql.WebSite, error) {
	panic("not implemented")
}
func (r *queryResolver) Documents(ctx context.Context) ([]*go_graphql.Document, error) {
	panic("not implemented")
}
func (r *queryResolver) Categories(ctx context.Context) ([]*go_graphql.Category, error) {
	panic("not implemented")
}
