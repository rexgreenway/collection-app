package storage

import "time"

// Utils holds overridable utility functions.
type Utils struct {
	// NowUTC returns the current time UTC.
	NowUTC func() time.Time
}

// Option functions can override service utils.
//
// Functions should be named `With<utility-param>Func`
type Option func(*Utils)

func WithNowUTCFunc(fn func() time.Time) Option {
	return func(u *Utils) { u.NowUTC = fn }
}
