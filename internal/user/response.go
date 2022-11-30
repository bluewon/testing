package user

type LoginResponse struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
}
