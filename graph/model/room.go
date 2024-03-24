package model

type Room struct {
	CreatedAt string
	ID        int
	Kind      RoomKindEnum
	Props     RoomProps `bun:"-"`
	UpdatedAt string
	Messages  []Message `bun:"-"`
}

type RoomProps struct {
	RoomID int
	Name   string
}

func (p RoomProps) GetID() int {
	return p.RoomID
}
