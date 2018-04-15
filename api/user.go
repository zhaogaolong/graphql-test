package api

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (r *Resolver) GetUser(ctx context.Context, arg struct{ ID int32 }) (*User, error) {
	return db.getUser(ctx, int32(arg.ID))
}

// func (r *Resolver) CreateUser(arg struct{ name string }) (*User, error) {
// 	var user User
// 	user.name = arg.name
// 	user.id = 2
// 	return &user, nil
// }

func (r *Resolver) Hello() *string {
	data := "Hello world!"
	return &data
}

// func (r *Resolver) Hello(ctx context.Context) (string, error) {
// 	return "Hello world!", nil
// }

// func (u *User) ID() *int32 {
// 	r := int32(u.id)
// 	return &r
// }
func (u *User) NAME() *string {
	return &u.Name
}

func (u *User) ID(ctx context.Context) *graphql.ID {
	userId := graphql.ID(u.Model.ID)
	return &userId

}
