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
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Println("Create a new USER")
	uuidValue := uuid.UUIDv4()

	user := &model.User{ID: uuidValue, FirstName: input.FirstName, LastName: input.LastName}
	users = append(users, user)
	return user, nil
}

// RemoveUser is the resolver for the removeUser field.
func (r *mutationResolver) RemoveUser(ctx context.Context, input model.DeleteUser) (*model.User, error) {
	index := -1
	for i, user := range users {
		if user.ID == input.ID {
			index = i
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
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	index := -1
	u := &model.User{}
	for i, user := range users {
		if user.ID == input.ID {
			index = i
			u = user
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
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, userNotFound
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var users []*model.User

func init() {
	log.Println("Init - Users array to be created")
	users = make([]*model.User, 0)
	users = append(users, &model.User{ID: uuid.UUIDv4(), FirstName: "Kevin", LastName: "Smith", Dob: "12/27/1987"})
	users = append(users, &model.User{ID: uuid.UUIDv4(), FirstName: "Lara", LastName: "Johnson", Dob: "11/7/1988"})
	log.Println("Init - Users array has been created")
}
