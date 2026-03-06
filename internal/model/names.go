package model

import "time"

type Name struct {
	DateOfCreation time.Time
	Name           string
}

func (n Name) HasExpired() bool {
	if time.Since(n.DateOfCreation) > 24*time.Hour {
		return true
	}
	return false
}
