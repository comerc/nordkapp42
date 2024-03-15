package model

type Message struct {
	CreatedAt string `json:"createdAt"`
	ID        int    `json:"id"`
	IsRead    bool   `json:"isRead"`
	MemberID  int    `json:"memberId"`
	RoomID    int    `json:"roomId"`
	Text      string `json:"text"`
	UpdatedAt string `json:"updatedAt"`
	// Member    *Member `json:"member,omitempty"`
}
