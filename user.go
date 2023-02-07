package app

type User struct {
        Id       int `json:"-" db:"id"`
        Role     string `json:"role" binding:"required"`
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
}