package model

type User struct {
	UserID       string `json:"userID"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Org          string `json:"org"`
	Role         string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Org      string `json:"org" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type AuthResponse struct {
	Token  string `json:"token"`
	UserID string `json:"userID"`
	Role   string `json:"role"`
}
