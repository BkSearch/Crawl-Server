package common

import "time"

type Question struct {
	ID           int       `db:"id, pk"`
	Content      string    `db:"content"`
	AmountAnswer int       `db:"amount_answer"`
	URL          string    `db:"URL"`
	PostAt       time.Time `db:"post_at"`
	Vote         int       `db:"vote"`
}
