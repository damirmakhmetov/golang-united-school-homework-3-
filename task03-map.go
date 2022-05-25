package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]string, 0, len(input))
	for _, k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	result = keys
	return result
}
