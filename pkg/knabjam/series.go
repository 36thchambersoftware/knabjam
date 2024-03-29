package knabjam

import (
	"time"
)

type Series struct {
	Title          string
	Description    string
	Type           string // Limited|Infinite
	Creator        User
	ReleaseDate    time.Time // The date it became available to collect
	OutOfPrintDate time.Time // The date packs will no longer be offered
	Cards          []Card
}
