package tri

// If returns `t` if `condition` is true, otherwise it returns `f`.
func If[T any](condition bool, t T, f T) T {
	if condition {
		return t
	}
	return f
}

// IfExec executes `t` if the `condition` is true, otherwise it executes `f`.
func IfExec[T any](condition bool, t func() T, f func() T) T {
	if condition {
		return t()
	}
	return f()
}

// IfExecWithErr executes `t` and returns its result and error if the `condition` is true,
// otherwise it executes `f` and returns its result and error.
func IfExecWithErr[T any](condition bool, t func() (T, error), f func() (T, error)) (T, error) {
	if condition {
		return t()
	}
	return f()
}

// IfExecOrVal executes `t` and returns its result if the `condition` is true,
// otherwise it returns the provided value `f`.
func IfExecOrVal[T any](condition bool, t func() T, f T) T {
	if condition {
		return t()
	}
	return f
}

// IfExecOrValWithErr executes `t` and returns its result and a nil error if the `condition` is true,
// otherwise it returns the provided value `f` and a nil error.
func IfExecOrValWithErr[T any](condition bool, t func() (T, error), f T) (T, error) {
	if condition {
		return t()
	}
	return f, nil
}

// IfValOrExec returns the provided value 't' if the `condition` is true,
// otherwise it executes `f` and returns its result.
func IfValOrExec[T any](condition bool, t T, f func() T) T {
	if condition {
		return t
	}
	return f()
}

// IfValOrExecWithErr returns the provided value `t` and a nil error if the `condition` is true,
// otherwise it executes `f` and returns its result and error.
func IfValOrExecWithErr[T any](condition bool, t T, f func() (T, error)) (T, error) {
	if condition {
		return t, nil
	}
	return f()
}
