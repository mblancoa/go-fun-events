package tools

func FromPointerArray[T any](arr []*T) []T {
	n := len(arr)
	newArr := make([]T, n)
	for i, _ := range arr {
		newArr[i] = *arr[i]
	}
	return newArr
}

func ToPointerArray[T any](arr []T) []*T {
	n := len(arr)
	newArr := make([]*T, n)
	for i, _ := range arr {
		newArr[i] = &arr[i]
	}
	return newArr
}
