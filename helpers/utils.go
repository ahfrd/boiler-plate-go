package helpers

import (
	"math"
	"strconv"
)

func PaginationHelpers(countData string, pageNumber string, size string) map[string]interface{} {
	floTotalRow, _ := strconv.ParseFloat(countData, 64)
	floRecordPerPage, _ := strconv.ParseFloat(size, 64)
	floCurrentPage, _ := strconv.ParseFloat(pageNumber, 64)
	totalPage := math.Ceil(floTotalRow / floRecordPerPage)
	currentPage := pageNumber
	firstRecord := (floCurrentPage - 1) * floRecordPerPage
	startRecord := firstRecord + 1
	pagination := map[string]interface{}{
		"total_record":    countData,
		"total_page":      strconv.FormatFloat(float64(totalPage), 'f', 0, 64),
		"record_per_page": size,
		"current_page":    currentPage,
		"start_record":    strconv.FormatFloat(startRecord, 'f', 0, 64),
		"first_record":    strconv.FormatFloat(firstRecord, 'f', 0, 64),
	}
	return pagination
}
