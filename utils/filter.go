package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func GetFilter(temp_filters map[string]string, db *gorm.DB) *gorm.DB {
	filters := db
	for key, filter := range temp_filters {
		if filter != "" {
			temp_filter := fmt.Sprint("%" + filter + "%")
			filters = filters.Where(key+" like ?", temp_filter)
		}
	}

	return filters

}

func GetLimitOffset(size int, page int) (limit int, offset int) {
	if size == 0 && page == 0 {
		limit = 10
		offset = 0
	} else if size == 0 && page != 0 {
		limit = 10
		offset = (page - 1) * limit
	} else if size != 0 && page == 0 {
		limit = size
		offset = 0
	} else {
		limit = size
		offset = (page - 1) * size
	}

	return limit, offset
}