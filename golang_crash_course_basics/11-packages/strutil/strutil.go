package strutil

// functions with uppercase name automatically get exported

func Reverse(myString string) string {
	// Rune literals are just 32-bit integer values (however they're untyped constants, so their type can change).
	// They represent unicode codepoints. For example, the rune literal 'a' is actually the number 97.
	runes := []rune(myString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
