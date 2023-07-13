package util_test

import (
	"testing"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/util"
	"github.com/stretchr/testify/assert"
)

// TestString2Int64Slice with empty string.
func TestString2Int64Slice_EmptyString(t *testing.T) {
	// Given an empty string
	s := ""
	// When call String2Int64Slice
	ids, err := util.String2Int64Slice(s)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, ids)
}

// TestString2Int64Slice with a character string.
func TestString2Int64Slice_CharacterString(t *testing.T) {
	// Given a character string
	s := "a"
	// When call String2Int64Slice
	ids, err := util.String2Int64Slice(s)
	// Then return an error
	assert.NotNil(t, err)
	assert.Empty(t, ids)
}

// TestString2Int64Slice with a string with integers separated by commas.
func TestString2Int64Slice_IntegersSeparatedByCommas(t *testing.T) {
	// Given a string with integers separated by commas
	s := "1,2,3"
	// When call String2Int64Slice
	ids, err := util.String2Int64Slice(s)
	// Then return a slice of integers
	assert.Nil(t, err)
	assert.Equal(t, []int64{1, 2, 3}, ids)
}

// TestValidateDate with empty string.
func TestValidateDate_EmptyString(t *testing.T) {
	// Given an empty string
	s := ""
	// When call ValidateDate
	err := util.ValidateDate(s)
	// Then return an error
	assert.NotNil(t, err)
}

// TestValidateDate with a non valid date string.
func TestValidateDate_NonValidDateString(t *testing.T) {
	// Given a non valid date string
	s := "Hello Bia Energy"
	// When call ValidateDate
	err := util.ValidateDate(s)
	// Then return an error
	assert.NotNil(t, err)
}

// TestValidateDate with a valid date string.
func TestValidateDate_ValidDateString(t *testing.T) {
	// Given a valid date string
	s := "2021-01-01"
	// When call ValidateDate
	err := util.ValidateDate(s)
	// Then return an error
	assert.Nil(t, err)
}
