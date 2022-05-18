package reverse_string

func ReverseString(input string) (output string) {
	var result []rune
	arr := []rune(input)

	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}

	return string(result)
}
