package app

import "time"

type Topic struct {
	Id string
	Tag string
	Skipped uint32
	Completed uint32
	Skippable bool
	LastRecall time.Time
}

