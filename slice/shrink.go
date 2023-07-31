package slice

func calCapacity(c int, l int) (int, bool) {
	if c < 256 {
		return c, false
	}
	if c > 1024 && (c/l) >= 2 {
		factor := 0.75
		return int(float32(c) * float32(factor)), true
	}
	if c <= 1024 && (c/l) >= 3 {
		return c / 2, true
	}
	return c, false
}

func Shrink[T any](src []T) []T {
	oldCap, length := cap(src), len(src)
	newCap, changed := calCapacity(oldCap, length)
	if !changed {
		return src
	}
	newSlice := make([]T, 0, newCap)
	newSlice = append(newSlice, src...)
	return newSlice
}
