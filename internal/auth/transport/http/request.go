package auth_http

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
