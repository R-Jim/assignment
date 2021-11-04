package dto

type CreateTweetRequestBody struct {
	Content string `validate:"gt=0,lte=50"`
}
