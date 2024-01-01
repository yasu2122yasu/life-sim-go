package request

type CheckLoginRequest struct {
	Email     string `json:"email"`
	AuthToken string `json:"auth-token"`
}
