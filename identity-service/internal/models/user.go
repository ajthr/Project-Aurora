package models

import "time"

type User struct {
	Id               int
	Name             string
	Email            string
	IsSignupComplete bool
	CreatedAt        time.Time
}
