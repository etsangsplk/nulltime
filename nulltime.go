// Package NullTime provides a time.Time that may be null.
// Copyright (c) 2011-2013, 'pq' Contributors
// Portions Copyright (C) 2011 Blake Mizerany
// Modified by Kirill Danshin <k@geek.gq> (c) 2016
package NullTime

import "time"

// NullTime represents a time.Time that may be null. NullTime implements the
// sql.Scanner interface so it can be used as a scan destination, similar to
// sql.NullString.
type NullTime struct {
	Time time.Time
	// Valid is true if Time is not NULL
	Valid bool
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value returns time value, isValid and error object
func (nt NullTime) Value() (time.Time, bool, error) {
	if !nt.Valid {
		return time.Time{}, false, nil
	}
	return nt.Time, true, nil
}
