package request

type InitialLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
