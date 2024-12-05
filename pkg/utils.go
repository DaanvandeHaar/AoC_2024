package pkg

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func Contains[T comparable](comparable T, array []T) bool {
	for _, item := range array {
		if item == comparable {
			return true
		}
	}

	return false
}

func ContainsAny[T comparable](set []T, contains []T) bool {
	lookup := make(map[T]bool)

	for _, item := range set {
		lookup[item] = true
	}

	for _, item := range contains {
		if _, exists := lookup[item]; exists {
			return true
		}
	}

	return false
}
