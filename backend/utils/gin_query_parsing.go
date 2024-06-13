package services

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNullableUintQueryValue(c *gin.Context, queryName string) *uint {
	queryValue := c.Query(queryName)
	if queryValue == "" {
		return nil
	}

	value, err := strconv.ParseUint(queryValue, 10, 64)
	if err != nil {
		return nil
	}

	uintValue := uint(value)
	return &uintValue
}
