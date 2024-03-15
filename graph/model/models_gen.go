// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Query struct {
}

type Subscription struct {
}

type RoomKindEnum string

const (
	RoomKindEnumChat           RoomKindEnum = "CHAT"
	RoomKindEnumPrivateChannel RoomKindEnum = "PRIVATE_CHANNEL"
	RoomKindEnumPrivateGroup   RoomKindEnum = "PRIVATE_GROUP"
	RoomKindEnumPublicChannel  RoomKindEnum = "PUBLIC_CHANNEL"
	RoomKindEnumPublicGroup    RoomKindEnum = "PUBLIC_GROUP"
)

var AllRoomKindEnum = []RoomKindEnum{
	RoomKindEnumChat,
	RoomKindEnumPrivateChannel,
	RoomKindEnumPrivateGroup,
	RoomKindEnumPublicChannel,
	RoomKindEnumPublicGroup,
}

func (e RoomKindEnum) IsValid() bool {
	switch e {
	case RoomKindEnumChat, RoomKindEnumPrivateChannel, RoomKindEnumPrivateGroup, RoomKindEnumPublicChannel, RoomKindEnumPublicGroup:
		return true
	}
	return false
}

func (e RoomKindEnum) String() string {
	return string(e)
}

func (e *RoomKindEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoomKindEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoomKindEnum", str)
	}
	return nil
}

func (e RoomKindEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
