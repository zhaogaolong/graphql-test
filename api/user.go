package api

import (
	"context"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (r *Resolver) GetUser(ctx context.Context, args struct{ Id int32 }) (*User, error) {
	return db.getUser(ctx, int32(args.Id))
}

func (r *Resolver) GetUsers(ctx context.Context) (*[]*User, error) {
	return db.getUsers(ctx)
}

func (r *Resolver) CreateUser(ctx context.Context, args struct{ Name string }) (*User, error) {
	return db.createUser(ctx, args.Name)
}

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
func (u *User) NAME(ctx context.Context) *string {
	return &u.Name
}

func (u *User) ID(ctx context.Context) *int32 {
	// return &u.Model.ID
	// fmt.Println(u.Model.ID)
	userId := int32(u.Model.ID)
	return &userId
	// userId := graphql.ID(base64.StdEncoding.EncodeToString([]byte(string(u.Model.ID))))
	// return &userId

	// x := graphql.ID("dfsdfsdf")
	// return &x
}
