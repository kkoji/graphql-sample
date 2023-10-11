package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/kkoji/gqlgen-todos/graph/generated"
	"github.com/kkoji/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.createUser(input)
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "TODO-3",
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "name",
		},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.getUsers()
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "TODO-1",
			Text: "My Todo 1",
			User: &model.User{
				ID:   "User-1",
				Name: "hsaki",
			},
			Done: true,
		},
		{
			ID:   "TODO-2",
			Text: "My Todo 2",
			User: &model.User{
				ID:   "User-1",
				Name: "hsaki",
			},
			Done: false,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
