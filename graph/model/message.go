package model

type Message struct {
	CreatedAt string
	ID        int
	IsRead    bool
	MemberID  int
	RoomID    int
	Text      string
	UpdatedAt string
	Member    *Member `bun:"-"`
}

func (m Message) GetID() int {
	return m.ID
}
