package dto

type LoginRequestBody struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
