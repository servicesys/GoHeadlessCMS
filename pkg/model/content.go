package model

import "time"

type Content struct {
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	Type       Type      `json:"type"`
	Category   Category  `json:"category"`
	Content    []byte    `json:"content"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Status     string    `json:"status"`
}
