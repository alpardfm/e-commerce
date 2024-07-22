package entity

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthLoginDashboardBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type AuthLoginDashboardHeader struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type AuthLoginDashboardResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type TokenLoginDashboardClaims struct {
	UID    string `json:"uid,omitempty"`
	Email  string `json:"email,omitempty"`
	RoleID string `json:"role_id,omitempty"`
	jwt.StandardClaims
}
