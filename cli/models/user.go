package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int64
	Phone string
}
type UserRepo struct {
	db *gorm.DB
}

// FindOne
func (u *UserRepo) FindOne(id int64) (*User, error) {
}
func (u *UserRepo) FindByPhone(phone string) (*User, error) {
}

// FindManyByA...
// UpdateOne
// UpdateMany
// DeleteOne
// DeleteMany
// Insert
