package helper

import "reflect"

func Paginate[T any](slice []T, page, pageSize int) []T {
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	// Check if startIndex is within bounds
	if startIndex >= len(slice) {
		return []T{} // Return an empty slice if the startIndex is out of bounds
	}

	// Check if endIndex is within bounds
	if endIndex > len(slice) {
		endIndex = len(slice)
	}

	// Use reflection to create a slice of the same type as the input
	resultSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice).Elem()), endIndex-startIndex, endIndex-startIndex).Interface()

	// Copy the elements to the result slice using reflection
	for i := 0; i < endIndex-startIndex; i++ {
		reflect.ValueOf(resultSlice).Index(i).Set(reflect.ValueOf(slice[startIndex+i]))
	}

	return resultSlice.([]T)
}
