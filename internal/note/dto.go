package note

type CreateNoteDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type UpdateNoteDTO struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	IsDone  bool   `json:"is_done,omitempty"`
}
