package model

import "time"

type Content struct {
	Uuid          string    `json:"uuid"`
	ContentType   Type      `json:"content_type"`
	Value         string    `json:"value"`
	CreatedOn     time.Time `json:"created_on"`
	ModifiedOn    time.Time `json:"modified_on"`
	ContentStatus string    `json:"content_status"`
}
