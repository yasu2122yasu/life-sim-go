package response

type InitialLoginResponse struct {
	Email     string `json:"email"`
	AuthToken string `json:"auth-token"`
}
