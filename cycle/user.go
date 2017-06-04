package cycle

import (	
	"time"
)

//for storage the user information authentication
type AuthInfo struct{
	ClientId string
	SecretKey string
}

type User struct {
	Name string
	Username string
	Password string //temp... we should search for a real alternative in web scenario
	Birth_day time.Time
	Email string
	Sex byte
}

type SimpleUser struct{
	User User
	Connectivity int
	AuthSpotify AuthInfo
	AuthDeezer AuthInfo
}

type AdminUser struct {
	AdminUser User 
}