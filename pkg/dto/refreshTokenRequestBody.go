package dto

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt"`
}
