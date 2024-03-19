package models

import (
	pb "grpc/user"
	"server/tools"
	"time"
)

type User struct {
	ID        uint32     `gorm:"primaryKey,omitempty"`
	Name      string    `gorm:"name,omitempty"`
	Password  string    `gorm:"password,omitempty"`
	Dni       string    `gorm:"dni,omitempty"`
	Phone     string    `gorm:"phone,omitempty"`
	Website   string    `gorm:"website,omitempty"`
	Address   string    `gorm:"address,omitepty"`
	CreatedAt time.Time `gorm:"created_ad,omitempty"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty"`
}

func (u *User) CreateUser() error {
	password, err := tools.GenerateHasKey(u.Password)
	if err != nil {
		return err
	}
	u.Password = password 
	if err := DB.Create(u).Error; err != nil {
		go loggerError.Printf("Operation CreateUser => %s", err)
		return err
	}
	return nil
}

func (u *User) UpdateUser() error {
	if u.Password != "" && len(u.Password) >= 8 && len(u.Password) < 12 {
		password, err := tools.GenerateHasKey(u.Password)
		if err != nil {
			return err
		}
		u.Password = password
	}
	if err := DB.Updates(u).Error; err != nil {
		go loggerError.Printf("Operation UpdateUser => %s", err)
		return err
	}
	return nil
}

func (u User) DeleteUser() error {
	if err := DB.Delete(u).Error; err != nil {
		go loggerError.Printf("Operation DeleteUser => %s", err)
		return err
	}
	return nil
}

func (u User) GetUsers() ([]*pb.User, error) {
	var (
		us []User
		pbU []*pb.User
		p = new(pb.User)
	)

	if err := DB.Find(&us).Error; err != nil {
		go loggerError.Printf("Operation GetUsers => %s", err)
		return nil, err
	}

	for _, i := range us {
		p.Address = i.Address
		p.Dni = i.Dni
		p.Id = i.ID
		p.Phone = i.Phone
		p.Name = i.Name
		p.Website = i.Website
		pbU = append(pbU, p)
	}

	return pbU, nil
}
