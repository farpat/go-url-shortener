package requests

type StoreUrlRequest struct {
	Url  string `json:"url" validate:"required,url,startswith=https://"`
	Slug string `json:"slug" validate:"required,unique_slug"`
}
