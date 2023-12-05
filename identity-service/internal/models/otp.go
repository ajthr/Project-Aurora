package models

import "time"

type OTP struct {
	Id         int
	Email      string
	Value      string
	Expiration time.Time
}
