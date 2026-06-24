package utilities

// Optional returns the value referenced by val if val is not nil, otherwise returns the zero value of T.
func Optional[T any](val *T) T {
	var zero T
	if val != nil {
		return *val
	}

	return zero
}
