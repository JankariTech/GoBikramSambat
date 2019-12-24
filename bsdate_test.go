package bsdate

import (
	"github.com/magiconair/properties/assert"
	"strconv"
	"testing"
)

type TestDateStruc struct {
	day   int
	month int
	year  int
}

var validDates = []TestDateStruc{
	{1, 1, 1970},
	{1, 1, 2076},
	{3, 2, 2074},
	{30, 12, 2100},
}

var invalidDates = []TestDateStruc{
	{1, 1, 0000},
	{0, 2, 2074},
	{10, 0, 2074},
	{000, 2, 2074},
	{-1, 2, 2074},
	{1, -1, 2074},
	{1, 2, -1},
	{33, 2, 2074},
	{1, 13, 2074},
	{1, 1, 1969}, //no data before BS 1970
	{1, 1, 2101}, //no data after BS 2100
	{1, 32, 2076}, //this month has only 31 days
	{12, 31, 2067}, //this month has only 30 days
}

func TestValidBSDates(t *testing.T) {
	for _, testCase := range validDates {
		t.Run(strconv.Itoa(testCase.year)+"-"+strconv.Itoa(testCase.month)+"-"+strconv.Itoa(testCase.day), func(t *testing.T) {
			nepaliDate, err := New(testCase.day, testCase.month, testCase.year)
			assert.Equal(t, err, nil)
			assert.Equal(t, nepaliDate.GetDay(), testCase.day)
			assert.Equal(t, nepaliDate.GetMonth(), testCase.month)
			assert.Equal(t, nepaliDate.GetYear(), testCase.year)
		})
	}
}

func TestInvalidBSDates(t *testing.T) {
	for _, testCase := range invalidDates {
		t.Run(strconv.Itoa(testCase.year)+"-"+strconv.Itoa(testCase.month)+"-"+strconv.Itoa(testCase.day), func(t *testing.T) {
			nepaliDate, err := New(testCase.day, testCase.month, testCase.year)
			assert.Equal(t, err.Error(), "not a valid date")
			assert.Equal(t, nepaliDate, nil)
		})
	}
}
