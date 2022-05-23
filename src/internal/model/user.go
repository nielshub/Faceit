package model

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Merge(other User) {
	if other.FirstName != "" {
		u.FirstName = other.FirstName
	}
	if other.LastName != "" {
		u.LastName = other.LastName
	}
	if other.Nickname != "" {
		u.Nickname = other.Nickname
	}
	if other.Password != "" {
		u.Password = other.Password
	}
	if other.Email != "" {
		u.Email = other.Email
	}
	if other.Country != "" {
		u.Country = other.Country
	}
	u.UpdatedAt = time.Now()
}
