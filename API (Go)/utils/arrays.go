package utils

func Find(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func FindAndDelete(array []string, item string) []string {
	index := 0
	for _, i := range array {
		if i != item {
			array[index] = i
			index++
		}
	}

	return array[:index]
}