package model

import "time"

type ContentValue struct {
	Uuid	string	`json:"uuid"`
	ContentTypeCod	string	`json:"content_type_cod"`
	Value	string	`json:"value"`
	CreatedOn	time.Time	`json:"created_on"`
	ModifiedOn	time.Time	`json:"modified_on"`
	ContentStatus	string	`json:"content_status"`
}