package models

type OTP struct {
	UserId     int64
	Value      string
	Expiration string
}
