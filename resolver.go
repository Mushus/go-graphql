package go_graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateWebSite(ctx context.Context, input NewWebSite) (*WebSite, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateDocumument(ctx context.Context, input NewDocumument) (*Document, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCategory(ctx context.Context, input NewCategory) (*Category, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) WebSites(ctx context.Context) ([]*WebSite, error) {
	panic("not implemented")
}
func (r *queryResolver) Documents(ctx context.Context) ([]*Document, error) {
	panic("not implemented")
}
func (r *queryResolver) Categories(ctx context.Context) ([]*Category, error) {
	panic("not implemented")
}
