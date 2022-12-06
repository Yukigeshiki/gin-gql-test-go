package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"

	"github.com/Yukigeshiki/go-gin-gql-test/graph/generated"
	"github.com/Yukigeshiki/go-gin-gql-test/graph/model"
	"github.com/dgryski/trifles/uuid"
)

var userNotFound = errors.New("cannot find user you are looking for")

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(_ context.Context, input model.NewUser) (*model.User, error) {
	log.Println("Create a new USER")
	uuidValue := uuid.UUIDv4()

	user := &model.User{ID: uuidValue, FirstName: input.FirstName, LastName: input.LastName}
	users = append(users, user)
	return user, nil
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(_ context.Context, input model.DeleteUser) (*model.User, error) {
	index := -1
	for i, user := range users {
		if user.ID == input.ID {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, userNotFound
	}
	user := users[index]
	users = append(users[:index], users[index+1:]...)

	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(_ context.Context, input model.UpdateUser) (*model.User, error) {
	index := -1
	u := &model.User{}
	for i, user := range users {
		if user.ID == input.ID {
			index = i
			u = user
			break
		}
	}
	if index == -1 {
		return nil, userNotFound
	}

	if input.FirstName != nil {
		u.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		u.LastName = *input.LastName
	}
	if input.Dob != nil {
		u.Dob = *input.Dob
	}
	users[index] = u

	return u, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(_ context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, userNotFound
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(_ context.Context) ([]*model.User, error) {
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
