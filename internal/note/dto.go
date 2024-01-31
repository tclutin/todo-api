package note

type CreateNoteDTO struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (dto CreateNoteDTO) Validate() bool {
	if dto.Name == "" || dto.Content == "" {
		return false
	}
	return true
}

type UpdateNoteDTO struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	IsDone  bool   `json:"is_done,omitempty"`
}
