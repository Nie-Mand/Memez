package store

type Meme struct {
	ID       string    `json:"id"`
	Content string `json:"content"`
	Rate   string    `json:"rate"`
}