package slice

func DeleteByIndex[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, IndexOutOfRange
	}
	res := src[index]
	if index == length-1 {
		src = src[:length-1]
		return src, res, nil
	}
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	return src, res, nil
}
