package data

import "time"

type Blog struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Poster     string    `json:"poster"`
	Content    string    `json:"content"`
	Author     int       `json:"author"`
	Upvotes    int       `json:"upvotes"`
	Downvotes  int       `json:"downvotes"`
	CategoryID int       `json:"category_id"`
	SubjectID  int       `json:"subject_id"`
	TopicID    int       `json:"topic_id"`
	CreatedAt  time.Time `json:"-"`
}
