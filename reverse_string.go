package reverse_string

var str string = "Hello there"

var _ = ReverseString(str)

func ReverseString(input string) (output string) {

	for _, v := range input {
		output = string(v) + output
	}

	return output
}
