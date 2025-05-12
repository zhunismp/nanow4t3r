package helpers

// TODO: Move this to a separate package
func WithFallback[T any](opt *T, fallback T) T {
	if opt != nil {
		return *opt
	}
	return fallback
}
