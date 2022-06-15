package utility

import (
	"fmt"
)

func GetPaginateURL(path string, page *int, limit *int, totalData int) (string, string, int) {
	var tempPage int = *page
	var tempLimit int = *limit

	if tempPage < 0 {
		tempPage = 1
	}

	if tempLimit < 0 {
		tempLimit = 10
	} else if tempLimit > 25 {
		tempLimit = 25
	}

	totalPages := 0
	if totalPages = totalData / tempLimit; totalData % tempLimit != 0 {
		totalPages += 1
	}

	if (tempPage > totalPages) {
		tempPage = totalPages
	}

	nextPage := fmt.Sprintf("%s?page=%d&limit=%d", path, tempPage+1, tempLimit)
	prevPage := fmt.Sprintf("%s?page=%d&limit=%d", path, tempPage-1, tempLimit)

	if (tempPage+1) > totalPages {
		nextPage = ""
		tempPage = totalPages
	} else if (tempPage-1) < 1 {
		nextPage = ""
		tempPage = 1
	}

	if tempPage >= 1 && tempLimit >= totalData {
		tempPage = 1
		tempLimit = totalData
		prevPage = ""
	}

	*page = tempPage
	*limit = tempLimit
	return nextPage, prevPage, totalPages
}