package token

import "time"

type Maker interface {
	//createtoken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//verify token checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
