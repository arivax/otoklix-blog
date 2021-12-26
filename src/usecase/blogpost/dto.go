package blogpost

// BlogPostReq DTO
type BlogPostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// BlogPostRes DTO
type BlogPostRes struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Published_at string `json:"published_at"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
