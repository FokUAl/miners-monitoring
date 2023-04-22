package pkg

func IsContain(object string, container []string) bool {
	for _, elem := range container {
		if elem == object {
			return true
		}
	}

	return false
}
