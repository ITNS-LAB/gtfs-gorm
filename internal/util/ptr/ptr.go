package ptr

// Float64 は float64 値からポインタを返します。
func Float64(v float64) *float64 {
	return &v
}

// Int は int 値からポインタを返します。
func Int(v int) *int {
	return &v
}

// String は string 値からポインタを返します。
func String(v string) *string {
	return &v
}

// Bool は bool 値からポインタを返します。
func Bool(v bool) *bool {
	return &v
}

// Ptr は任意の型 T の値からポインタを返します。
// この関数を使用するには、Go 1.18 以降が必要です。
func Ptr[T any](v T) *T {
	return &v
}
