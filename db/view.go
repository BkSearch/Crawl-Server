package db

import (
  "time"
)

type questionAPIView struct {
	ID           int       `json:"id"`
	Content      string    `json:"content"`
	AmountAnswer int       `json:"amount_answer"`
	URL          string    `json:"URL"`
	PostAt       time.Time `json:"post_at"`
	Vote         int       `json:"vote"`
}
