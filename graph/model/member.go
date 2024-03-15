package model

type Member struct {
	CreatedAt string `json:"createdAt"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updatedAt"`
}

func (m Member) GetID() int {
	return m.ID
}
