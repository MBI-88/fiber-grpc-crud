package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey,omitempty"`
	Name      string    `gorm:"name,omitempty"`
	Password  string    `gorm:"password,omitempty"`
	Dni       string    `gorm:"dni,omitempty"`
	Phone     string    `gorm:"phone,omitempty"`
	Website   string    `gorm:"website,omitempty"`
	Address   string    `gorm:"address,omitepty"`
	CreatedAt time.Time `gorm:"created_ad,omitempty"`
}

func (u *User) CreateUser() error {
	return nil
}

func (u *User) UpdateUser() error {
	return nil
}

func (u User) DeleteUser() error {
	return nil
}

func (u User) GetUsers() ([]User, error) {
	var us []User
	return us, nil
}
