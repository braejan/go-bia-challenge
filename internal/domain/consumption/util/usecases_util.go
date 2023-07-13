package util

import (
	"strconv"
	"strings"
	"time"

	consumptionVO "github.com/braejan/go-bia-challenge/internal/valueobject/consumption"
)

// String2Int64Slice Convert string to int64 slice.
// Example: "1,2,3" -> []int64{1,2,3}
func String2Int64Slice(s string) (ids []int64, err error) {
	idsArray := strings.Split(s, ",")
	for _, id := range idsArray {
		idInt64, err := String2Int64(id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, idInt64)
	}
	return
}

// String2Int64 Convert string to int64.
// Example: "1" -> int64(1)
func String2Int64(s string) (id int64, err error) {
	// parse string to int64
	id, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return
}

// Validate date string in format YYYY-MM-DD
func ValidateDate(dateStr string) (err error) {
	// Check date is not empty
	if dateStr == "" {
		return consumptionVO.ErrEmptyDate
	}
	// convert string to time
	_, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		return consumptionVO.ErrInvalidDate
	}
	return
}
