package homework

func average(input [15]float32) (result float32) {
	i := 0
	var sum float32
	for i < len(input) {
		sum = input[i] + sum
	}
	result = sum / float32(len(input))
	return result
}
