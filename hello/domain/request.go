package domain

type LoginRequest struct {
	Username string
	Password string
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}

type PostRequest struct {
	Id     string `json:"id" param:"id"`
	Name   string `json:"name" form:"name"`
	Data   string `json:"data"`
	Value  string `json:"value"`
	Number uint64 `json:"number"`
}
