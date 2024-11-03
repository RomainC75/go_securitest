package db

import (
	"database/sql"
	"time"
)

type IpRange struct {
	ID       int64          `json:"id"`
	ScanID   int64          `json:"scanId"`
	IpMin    string         `json:"ipMin"`
	IpMax    sql.NullString `json:"ipMax"`
	IsUnique sql.NullBool   `json:"isUnique"`
}

type PortRange struct {
	ID       int64         `json:"id"`
	ScanID   int64         `json:"scanId"`
	RangeMin int32         `json:"rangeMin"`
	RangeMax sql.NullInt32 `json:"rangeMax"`
	IsUnique sql.NullBool  `json:"isUnique"`
}

type Scan struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Scenario  int32     `json:"scenario"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
