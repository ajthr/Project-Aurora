package models

// request body for signup
type SignUpRequest struct {
	Name  string
	Email string
}

// request body for signin
type SignInRequest struct {
	Name  string
	Email string
}

// request body for google signin
type GoogleSigninRequest struct {
	Token string
}

type OTPValidationRequest struct {
	Email string
	OTP   string
}
