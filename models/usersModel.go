package models

// Users Model
type User struct {
	Id           string `json:"_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	IsBanned     bool   `json:"is_banned"`
	ProfilePic   string `json:"profile_pic"`
	Country      string `json:"country"`
	LastLoggedIn string `json:"last_logged_in"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
