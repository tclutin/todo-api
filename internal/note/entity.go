package note

import "time"

type Note struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
