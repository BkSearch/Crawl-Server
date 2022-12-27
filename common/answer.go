package common

import "time"

type Answer struct {
	ID          int       `db:"id, pk"`
	Content     string    `db:"content"`
	URL         string    `db:"url"`
	PostAt      time.Time `db:"post_at"`
	Vote        int       `db:"vote"`
	Question_ID int       `db:"question_id"`
}
