package model

type AuthRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
