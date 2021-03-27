package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-gqlgen-graphql-sample/graph/generated"
	"go-gqlgen-graphql-sample/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var targetUser *model.User
	for _, user := range r.users {
		if user.ID == input.UserID {
			targetUser = user
			break
		}
	}
	if targetUser == nil {
		return nil, fmt.Errorf("user with id='%s' not found", input.UserID)
	}
	newTodo := &model.Todo{
		ID:   strconv.Itoa(r.lastTodoId),
		Text: input.Text,
		Done: false,
		User: targetUser,
	}
	r.todos = append(r.todos, newTodo)
	r.lastTodoId++
	return newTodo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
