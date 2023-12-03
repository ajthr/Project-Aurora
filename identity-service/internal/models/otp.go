package models

import "time"

type OTP struct {
	Id         int
	UserId     int
	Value      string
	Expiration time.Time
}
