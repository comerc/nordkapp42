package model

type Room struct {
	CreatedAt string
	ID        int
	Kind      RoomKindEnum
	Name      string
	UpdatedAt string
	Messages  []Message `bun:"-"`
}
