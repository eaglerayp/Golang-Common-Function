package main

import (
	"fmt"
	"time"
)

// Utilities function

// TimestampGen generate a unix millisecond.
func TimestampGen(changeTime time.Time) int64 {
	return changeTime.UnixNano() / int64(time.Millisecond)
}

// depth 1 map compare
func MapEqual(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

// Check ISODate String
func IsISODate(str string) bool {
	_, err := time.Parse(time.RFC3339, str)
	if err == nil {
		return true
	}
	_, err = time.Parse(time.RFC3339Nano, str)
	if err == nil {
		return true
	}
	_, err = time.Parse("2016-04-14T01:17:25.980+0000", str)
	if err != nil {
		return true
	}
	_, err = time.Parse("2006-01-02T15:04:05.999999999Z0700", str)
	if err != nil {
		return true
	}
	// Date case
	return false
}

func ISODateGen(timestamp int64) time.Time {
	nanoTimestamp := int64(time.Millisecond) * timestamp
	Date := time.Unix(0, nanoTimestamp)
	return Date
}

func ISOStringGen() string {
	t := time.Now()
	return t.UTC().Format(time.RFC3339)
}

func SetToStringSlice(PGIDSet map[string]bool) (PGIDStrings []string) {
	for k, _ := range PGIDSet {
		PGIDStrings = append(PGIDStrings, k)
	}
	return
}

func main() {
	fmt.Println(IsISODate("2012-11-01T22:08:41+00:00"))
	fmt.Println(IsISODate("2016-04-14T01:17:25.980+0000"))
	fmt.Println(IsISODate("2016-04-14T01:17:25.980Z"))
	fmt.Println(IsISODate("2012-11-01T22:08:41+00:00"))
}
