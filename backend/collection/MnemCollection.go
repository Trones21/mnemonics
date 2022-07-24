package collection

import "time"

type MnemCollection struct {
	CollectionID string
	CreatedByID  string
	CreatedDate  time.Time
	IsPublic     bool
}
