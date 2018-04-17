package mystring

// XReverse -- Starting with caps 'R' to indicate this function is exported,
// and can be used in other packages, when they import this package
func XReverse(s string) string {
	//b := make([]byte, len(s))
	b := make([]rune, len(s))
	j := 0
	for i := len(b) - 1; i >= 0; i-- {
		b[j] = rune(s[i]) //This will not work for all unicode chars. So operate on rune, instead of string
		j++
	}
	return string(b)
}

// Reverse -- Operates of rune and swap the contents over to generate reverse.
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[j], b[i] = b[i], b[j]
	}
	return string(b)
}
