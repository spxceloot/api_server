package types

type User struct {
	UID       string
	Email     string
	Password  string
	Username  string
	CreatedAt string
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
