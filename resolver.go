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

func (r *queryResolver) WebSite(ctx context.Context, id string) (*WebSite, error) {
	panic("not implemented")
}
func (r *queryResolver) Document(ctx context.Context, id string) (*Document, error) {
	panic("not implemented")
}
func (r *queryResolver) Category(ctx context.Context, id *string, name *string) (*Category, error) {
	panic("not implemented")
}
func (r *queryResolver) AllWebSites(ctx context.Context) ([]*WebSite, error) {
	panic("not implemented")
}
func (r *queryResolver) AllDocuments(ctx context.Context, input *QueryDocuments) ([]*Document, error) {
	panic("not implemented")
}
func (r *queryResolver) AllCategories(ctx context.Context, input *QueryCategories) ([]*Category, error) {
	panic("not implemented")
}
