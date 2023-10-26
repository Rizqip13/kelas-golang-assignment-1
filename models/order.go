package models

import (
	"time"
)

type Order struct {
	ID           uint       `gorm:"primaryKey" json:""`
	CustomerName string     `gorm:"not null" json:"customerName"`
	OrderedAt    *time.Time `json:"orderedAt"`
	Items        []Item     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

// func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
// 	now := time.Now()
// 	o.UpdatedAt = &now
// 	return
// }
