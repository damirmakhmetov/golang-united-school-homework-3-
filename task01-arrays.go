package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var k int = 0
	for i := range input {
		if input[i] == 0 {
			break
		}
		sum = input[i] + sum
		k += 1
	}
	result = sum / float32(k)
	return result
}
