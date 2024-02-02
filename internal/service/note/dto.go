package note

type CreateNoteDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type UpdateNoteDTO struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	IsDone  bool   `json:"is_done"`
}

func (dto CreateNoteDTO) Validate() bool {
	if dto.Name == "" || dto.Content == "" {
		return false
	}
	return true
}
