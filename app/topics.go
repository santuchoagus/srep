package app

import "time"

type Topic struct {
	Id         string
	Tag        string
	Skipped    int
	Completed  int
	Skippable  bool
	LastRecall time.Time
}
