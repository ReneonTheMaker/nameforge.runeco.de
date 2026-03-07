package model

import "time"

type RenderConfig struct {
	FileName       bool
	Lowercase      bool
	Project        bool
	VersionNumber  int
	FileExtension  string
	TimeOfCreation time.Time
}

func (c RenderConfig) HasExpired() bool {
	if time.Since(c.TimeOfCreation) > 24*time.Hour {
		return true
	}
	return false
}
