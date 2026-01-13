package internal

func MakeIncrementingIndexCounter() func() (int, error) {
	idx := -1
	return func() (int, error) {
		idx += 1
		return idx, nil
	}
}
// ID-1768294480-acc2ec9b
