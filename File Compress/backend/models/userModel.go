package models

import "time"

type User struct {
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
}
