package http

import "time"

type PostsRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	PictureURL string `json:"picture_url"`
	UserID     string `json:"user_id"`
}

type PostsResponse struct {
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	PictureURL string     `json:"picture_url"`
	UserID     string     `json:"user_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
