package helper

func Paginate[T any](slice []T, page, pageSize int) []T {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	startIndex := (page - 1) * pageSize
	if startIndex >= len(slice) {
		return []T{}
	}

	endIndex := startIndex + pageSize
	if endIndex > len(slice) {
		endIndex = len(slice)
	}

	return slice[startIndex:endIndex]
}
