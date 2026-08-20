package collection

// Utils holds overridable utility functions.
type Utils struct {
	// NewId generates a unique identifier for resources.
	NewId func() string
}

// Option functions can override service utils.
//
// Functions should be named `With<utility-param>Func`
type Option func(*Utils)

func WithNewIdFunc(fn func() string) Option {
	return func(u *Utils) { u.NewId = fn }
}
