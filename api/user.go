package api

type User struct {
	id   int32
	name string
}

func (r *Resolver) GetUser() (*User, error) {
	var user User
	user.name = "testUser"
	user.id = 1
	return &user, nil
}

func (r *Resolver) CreateUser(arg struct{ name string }) (*User, error) {
	var user User
	user.name = arg.name
	user.id = 2
	return &user, nil
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
func (u *User) NAME() *string {
	return &u.name
}

func (u *User) ID() *int32 {
	return &u.id
}
