package algorithm

func Transform[A any, B any](input []A, f func(A) (B, error)) ([]B, error) {
	results := make([]B, 0, len(input))

	for _, x := range input {
		y, err := f(x)
		if err != nil {
			return nil, err
		}
		results = append(results, y)
	}

	return results, nil
}
