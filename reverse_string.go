package reverse_string

func ReverseString(input string) (output string) {
	reverse := []rune(input)

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}
