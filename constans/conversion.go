package constans

// Conversion This func conversion type Url and Cors to type string
func Conversion[T Url | Cors](item T) string {
	return string(item)
}
