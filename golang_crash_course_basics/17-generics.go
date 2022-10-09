package main

import "fmt"

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T // create variable to fill it with its default type
		return zero, fmt.Errorf("empty")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}

	return m, nil
}

func main() {
	fmt.Println(min([]float64{3, 1, 2}))
	fmt.Println(min([]string{"Let's go", "Hey", "Ho"}))
}
