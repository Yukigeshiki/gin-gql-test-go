package graph

import (
	"github.com/Yukigeshiki/go-gin-gql-test/graph/model"
	"github.com/dgryski/trifles/uuid"
	"log"
)

var users []*model.User

func init() {
	log.Println("Init - Users array to be created")
	users = make([]*model.User, 0)
	users = append(users, &model.User{ID: uuid.UUIDv4(), FirstName: "Kevin", LastName: "Smith", Dob: "12/27/1987"})
	users = append(users, &model.User{ID: uuid.UUIDv4(), FirstName: "Lara", LastName: "Johnson", Dob: "11/7/1988"})
	log.Println("Init - Users array has been created")
}
