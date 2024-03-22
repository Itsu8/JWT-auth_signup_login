package modules

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 			uuid.UUID
	Created_at	time.Time
	Updated_at	time.Time
	Deleted_at  time.Time
	Username 	string  
	Password    string
	Age 		uint	
	Bio			string  
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) GetUserAge() uint {
	return u.Age
}

func (u User) GetUserBio() string {
	return u.Bio
}