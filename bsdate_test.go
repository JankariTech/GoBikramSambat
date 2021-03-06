package bsdate

import (
	"github.com/magiconair/properties/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

type TestDateStruc struct {
	day               int
	month             int
	expectedMonthName string
	year              int
}

type TestDateStrucWithMonthNames struct {
	day              int
	month            string
	expectedMonthNum int
	year             int
}

type TestDateConversionStruc struct {
	bsDate        string
	gregorianDate string
}

var validDates = []TestDateStruc{
	{1, 1, "Baisakh", 1970},
	{1, 1, "Baisakh", 2076},
	{3, 2, "Jestha", 2074},
	{30, 12, "Chaitra", 2100},
}

var validDatesWithMonthNames = []TestDateStrucWithMonthNames{
	{1, "Baisakh", 1, 1970},
	{2, "Jestha", 2, 1987},
	{1, "Ashadh", 3, 2076},
	{20, "Shrawan", 4, 1987},
	{21, "Bhadra", 5, 1987},
	{12, "Ashwin", 6, 2074},
	{13, "Kartik", 7, 2074},
	{17, "Mangsir", 8, 2074},
	{19, "Paush", 9, 2074},
	{22, "Mangh", 10, 2074},
	{15, "Falgun", 11, 2074},
	{30, "Chaitra", 12, 2100},
}

var invalidDates = []TestDateStruc{
	{1, 1, "", 0000},
	{0, 2, "", 2074},
	{10, 0, "", 2074},
	{000, 2, "", 2074},
	{-1, 2, "", 2074},
	{1, -1, "", 2074},
	{1, 2, "", -1},
	{33, 2, "", 2074},
	{1, 13, "", 2074},
	{1, 1, "", 1969},   //no data before BS 1970
	{1, 1, "", 2101},   //no data after BS 2100
	{32, 1, "", 2076},  //this month has only 31 days
	{31, 12, "", 2067}, //this month has only 30 days
	{30, 13, "", 2070},
}

var convertedDates = []TestDateConversionStruc{
	{"2068-04-01", "2011-07-17"}, //a random date
	{"2068-01-01", "2011-04-14"}, //1st Basakh
	{"2037-11-28", "1981-03-11"},
	{"2038-09-17", "1982-01-01"}, //1st Jan
	{"2040-09-17", "1984-01-01"}, //1st Jan in a leap year
	{"2040-09-18", "1984-01-02"}, //second Jan in a leap year
	{"2041-09-17", "1984-12-31"}, //31th Dec in a leap year
	{"2041-09-18", "1985-01-01"}, //1st Jan after a leap year
	{"2068-09-01", "2011-12-16"}, //1st Paush
	{"2068-08-29", "2011-12-15"}, //last day before first Paush
	{"2068-09-20", "2012-01-04"},
	{"2077-08-30", "2020-12-15"}, //last day before first Paush in a leap year
	{"2077-09-16", "2020-12-31"}, //31th Dec in a later leap year
	{"2074-09-16", "2017-12-31"}, //31th Dec in a non leap year
	{"2077-09-17", "2021-01-01"}, //1st Jan after a leap year
	{"2077-09-01", "2020-12-16"}, //1st Paush in a leap year
	{"2076-11-17", "2020-02-29"}, //29th Febr in a leap year
	{"2076-11-18", "2020-03-01"}, //1st March in a leap year
	{"2075-11-16", "2019-02-28"}, //28th Febr in a non leap year
	{"2076-02-01", "2019-05-15"}, //start of a month with 32 days
	{"2076-02-32", "2019-06-15"}, //end of a month with 32 days
	{"2076-03-01", "2019-06-16"}, //a month after a month with 32 days
	{"2100-12-30", "2044-04-12"}, //last day, we can convert in both directions
	{"2076-11-18", "2020-03-01"},
}
func TestValidBSDates(t *testing.T) {
	for _, testCase := range validDates {
		t.Run(strconv.Itoa(testCase.year)+"-"+strconv.Itoa(testCase.month)+"-"+strconv.Itoa(testCase.day), func(t *testing.T) {
			nepaliDate, err := New(testCase.day, testCase.month, testCase.year)
			assert.Equal(t, err, nil)
			assert.Equal(t, nepaliDate.GetDay(), testCase.day)
			assert.Equal(t, nepaliDate.GetMonth(), testCase.month)
			assert.Equal(t, nepaliDate.GetMonthName(), testCase.expectedMonthName)
			assert.Equal(t, nepaliDate.GetYear(), testCase.year)
		})
	}
}

func TestValidBSDatesWithMonthNames(t *testing.T) {
	for _, testCase := range validDatesWithMonthNames {
		t.Run(strconv.Itoa(testCase.year)+"-"+testCase.month+"-"+strconv.Itoa(testCase.day), func(t *testing.T) {
			nepaliDate, err := New(testCase.day, testCase.month, testCase.year)
			assert.Equal(t, err, nil)
			assert.Equal(t, nepaliDate.GetDay(), testCase.day)
			assert.Equal(t, nepaliDate.GetMonth(), testCase.expectedMonthNum)
			assert.Equal(t, nepaliDate.GetMonthName(), testCase.month)
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

func TestInvalidMonthName(t *testing.T) {
	nepaliDate, err := New(1, "NotExistingMonth", 2076)
	assert.Equal(t, err.Error(), "not a valid date")
	assert.Equal(t, nepaliDate, nil)
}

func TestInvalidMonthType(t *testing.T) {
	nepaliDate, err := New(1, 2.345, 2076)
	assert.Equal(t, err.Error(), "month has to be of value int or string")
	assert.Equal(t, nepaliDate, nil)
}

func TestConversionToGregorian(t *testing.T) {
	for _, testCase := range convertedDates {
		t.Run(testCase.bsDate, func(t *testing.T) {
			var bsYear, bsMonth, bsDay = splitDateString(testCase.bsDate)
			nepaliDate, err := New(bsDay, bsMonth, bsYear)
			assert.Equal(t, err, nil)

			var convertedGregorianDate, _ = nepaliDate.GetGregorianDate()
			expectedGregorianDate, _ := time.Parse("2006-01-02", testCase.gregorianDate)
			assert.Equal(t, convertedGregorianDate, expectedGregorianDate)
		})

	}
}

//cannot convert anything before BS 1970-09-01 because for that we would need data from 1969
var impossibleToConvertToGregorianDates = [] string {
	"1970-01-01",
	"1970-08-29",
}
func TestConversionInvalidToGregorian(t *testing.T) {
	for _, testCase := range impossibleToConvertToGregorianDates {
		t.Run(testCase, func(t *testing.T) {
			var bsYear, bsMonth, bsDay = splitDateString(testCase)
			nepaliDate, err := New(bsDay, bsMonth, bsYear)
			assert.Equal(t, err, nil)
			var convertedGregorianDate time.Time
			convertedGregorianDate, err = nepaliDate.GetGregorianDate()
			expectedGregorianDate, _ := time.Parse("2006-01-02", "0001-01-01")
			assert.Equal(t, err.Error(), "cannot convert date, invalid or missing data")
			assert.Equal(t, convertedGregorianDate, expectedGregorianDate)
		})
	}
}

func TestCreateFromGregorian(t *testing.T) {
	for _, testCase := range convertedDates {
		t.Run(testCase.bsDate, func(t *testing.T) {
			var expectedBsYear, expectedBsMonth, expectedBsDay = splitDateString(testCase.bsDate)
			var gregorianYear, gregorianMonth, gregorianDay = splitDateString(testCase.gregorianDate)
			nepaliDate, err := NewFromGregorian(gregorianDay, gregorianMonth, gregorianYear)
			assert.Equal(t, err, nil)
			assert.Equal(t, nepaliDate.GetDay(), expectedBsDay)
			assert.Equal(t, nepaliDate.GetMonth(), expectedBsMonth)
			assert.Equal(t, nepaliDate.GetYear(), expectedBsYear)
		})
	}
}

var impossibleToConvertFromGregorianDates = [] string {
	"1913-09-01", //1913 and before cannot be converted, because 1913+56=1969 and we do not have data for that BS year
	"1913-12-31",
	"2045-01-01", //2045 and after cannot be converted because 2045+56=2101 and we do not have data for that BS year
	"2045-05-01",
	"2044-04-13", //this date would be BS 2101-01-01  and we do not have data for that BS year
	"2019-02-29",
	"2019-13-22",
	"2010-11-32",
	"2020-02-30",
}
func TestCreateFromInvalidGregorian(t *testing.T) {
	for _, testCase := range impossibleToConvertFromGregorianDates {
		t.Run(testCase, func(t *testing.T) {

			var gregorianYear, gregorianMonth, gregorianDay = splitDateString(testCase)
			bsDate, err := NewFromGregorian(gregorianDay, gregorianMonth, gregorianYear)
			assert.Equal(t, err.Error(), "cannot convert date, invalid or missing data")
			assert.Equal(t, bsDate, nil)
		})
	}
}

func splitDateString(dateString string) (year int, month int, day int) {
	var splitedDate = strings.Split(dateString, "-")
	day, _ = strconv.Atoi(splitedDate[2])
	month, _ = strconv.Atoi(splitedDate[1])
	year, _ = strconv.Atoi(splitedDate[0])
	return
}
