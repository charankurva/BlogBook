package data

import "time"

type Blog struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Poster     string
	Content    string
	Author     int
	Upvotes    int
	Downvotes  int
	CategoryID int
	SubjectID  int
	TopicID    int
	CreatedAt  time.Time `json:"created_at"`
}
