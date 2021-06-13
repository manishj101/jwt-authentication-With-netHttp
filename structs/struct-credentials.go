package structs

// Response :
type Response struct {
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
}

type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	UserType string `json:"usertype,omitempty"`
	Email    string `json:"email"`
}
