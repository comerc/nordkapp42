package model

type Room struct {
	CreatedAt string       `json:"createdAt"`
	ID        int          `json:"id"`
	Kind      RoomKindEnum `json:"kind"`
	Name      *string      `json:"name,omitempty"`
	UpdatedAt string       `json:"updatedAt"`
	// Messages  []*Message
}
