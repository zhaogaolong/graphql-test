package api

import (
	"context"

	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var users = []User{
	User{Name: "zhaogaolong"},
	User{Name: "conny"},
	User{Name: "user1"},
}

// DB is the DB that will performs all operation
type DB struct {
	DB *gorm.DB
}

var db DB

func init() {
	var err error
	d, err := newDB("./db.sqlite")
	if err != nil {
		panic(err)
	}

	db = *d
}

func newDB(path string) (*DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
}

func (db *DB) getUser(ctx context.Context, id int32) (*User, error) {
	var user User

	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) getUsers(ctx context.Context) (*[]*User, error) {
	var users []*User

	err := db.DB.Limit(10).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (db *DB) createUser(ctx context.Context, name string) (*User, error) {
	user := User{
		Name: name,
	}

	err := db.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
