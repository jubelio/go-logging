package logging

func sContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func sIndex(s []string, str string) int {
	for i, v := range s {
		if v == str {
			return i
		}
	}

	return -1
}
