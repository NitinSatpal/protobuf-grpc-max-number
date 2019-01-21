package types

import (
	"time"
)

// User is a representation of a DB-stored user record.
type User struct {
	ID int `storm:"id,increment"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`

	Identity string `json:"identity"`
	Key      string `json:"key"`

	CurrentMaxNumber int
}

// InitTimes initializes time.Time variables in the header to current time.
func (h *User) InitTimes() {
	if h.CreatedAt.IsZero() {
		h.CreatedAt = time.Now()
	}
	h.UpdatedAt = h.CreatedAt
	h.DeletedAt = nil
}

func (h *User) IsDeleted() bool {
	return h.DeletedAt != nil
}
