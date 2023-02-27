package models

import "time"

type Base struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp; column:created_at; not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; column:updated_at; not null" json:"updated_at"`
}
