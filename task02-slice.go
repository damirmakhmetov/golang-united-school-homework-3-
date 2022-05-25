package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i := range input {
		result = append(result, input[len(input)-i-1])
	}
	return result
}
