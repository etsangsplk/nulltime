// Package NullTime provides a time.Time that may be null.
// Copyright (c) 2011-2013, 'pq' Contributors
// Portions Copyright (C) 2011 Blake Mizerany
package NullTime

import (
	"testing"
	"time"
)

func TestScanTimestamp(t *testing.T) {
	var nt NullTime
	tn := time.Now()
	nt.Scan(tn)
	if !nt.Valid {
		t.Errorf("Expected Valid=false")
	}
	if nt.Time != tn {
		t.Errorf("Time value mismatch")
	}
}

func TestValue(t *testing.T) {
	var (
		nt           NullTime
		returnedTime time.Time
		isValid      bool
	)
	returnedTime, isValid, _ = nt.Value()
	if isValid {
		t.Errorf("Expected isValid false but found true")
	}
	var emptyTime time.Time
	emptyTime = time.Time{}
	if returnedTime != emptyTime {
		t.Errorf("Time value mismatch")
	}
	tn := time.Now()
	nt.Scan(tn)
	returnedTime, isValid, _ = nt.Value()
	if !isValid {
		t.Errorf("Expected isValid true but found false")
	}
	if returnedTime != tn {
		t.Errorf("Time value mismatch")
	}

}

func TestScanNilTimestamp(t *testing.T) {
	var nt NullTime
	nt.Scan(nil)
	if nt.Valid {
		t.Errorf("Expected Valid=false")
	}
}
