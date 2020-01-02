package bsdate

import (
	"errors"
	"time"
)

type Date interface {
	GetDay() int
	GetMonth() int
	GetYear() int
	GetMonthName() string
	GetGregorianDate() (time.Time, error)
}
type date struct {
	Day        int
	Month      int
	Year       int
	MonthNames [12]string
}


// nepali year : day in Paush for 1st Jan, no of days in Baisakh, no of days in Jestha, no of days in Ashadh, ..
var calendardata = map[int][13]int {
	1970: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1971: [13]int{18, 31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	1972: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	1973: [13]int{19, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	1974: [13]int{19, 31, 31, 32, 30, 31, 31, 30, 29, 30, 29, 30, 30},
	1975: [13]int{18, 31, 31, 32, 32, 30, 31, 30, 29, 30, 29, 30, 30},
	1976: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	1977: [13]int{18, 31, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31},
	1978: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1979: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	1980: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	1981: [13]int{18, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	1982: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1983: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	1984: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	1985: [13]int{18, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	1986: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1987: [13]int{18, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	1988: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	1989: [13]int{18, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	1990: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1991: [13]int{18, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},

	//this data are from http://nepalicalendar.rat32.com/index.php
	1992: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	1993: [13]int{18, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	1994: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1995: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	1996: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	1997: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1998: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	1999: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2000: [13]int{17, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2001: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2002: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2003: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2004: [13]int{17, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2005: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2006: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2007: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2008: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2009: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2010: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2011: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2012: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2013: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2014: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2015: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2016: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2017: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2018: [13]int{18, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2019: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2020: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2021: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2022: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2023: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2024: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2025: [13]int{18, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2026: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2027: [13]int{17, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2028: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2029: [13]int{18, 31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2030: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 30, 30, 30, 31},
	2031: [13]int{17, 31, 32, 31, 32, 31, 31, 31, 31, 31, 31, 31, 31},
	2032: [13]int{17, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
	2033: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2034: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2035: [13]int{17, 30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2036: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2037: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2038: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2039: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2040: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2041: [13]int{18, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2042: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2043: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2044: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2045: [13]int{18, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2046: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2047: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2048: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2049: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2050: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2051: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2052: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2053: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2054: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2055: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 30, 29, 30},
	2056: [13]int{17, 31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
	2057: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2058: [13]int{17, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2059: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2060: [13]int{17, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2061: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2062: [13]int{17, 30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31},
	2063: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2064: [13]int{17, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2065: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2066: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
	2067: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2068: [13]int{17, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2069: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2070: [13]int{17, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
	2071: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2072: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2073: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
	2074: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2075: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2076: [13]int{16, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
	2077: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
	2078: [13]int{17, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
	2079: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
	2080: [13]int{16, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},

	//this data are from http://www.ashesh.com.np/nepali-calendar/
	2081: [13]int{17, 31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2082: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2083: [13]int{17, 31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2084: [13]int{17, 31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2085: [13]int{17, 31, 32, 31, 32, 31, 31, 30, 30, 29, 30, 30, 30},
	2086: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2087: [13]int{16, 31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30},
	2088: [13]int{16, 30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30},
	2089: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2090: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2091: [13]int{16, 31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30},
	2092: [13]int{16, 31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2093: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2094: [13]int{17, 31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
	2095: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 30, 30, 30, 30},
	2096: [13]int{17, 30, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
	2097: [13]int{17, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
	2098: [13]int{17, 31, 31, 32, 31, 31, 31, 29, 30, 29, 30, 30, 31},
	2099: [13]int{17, 31, 31, 32, 31, 31, 31, 30, 29, 29, 30, 30, 30},
	2100: [13]int{17, 31, 32, 31, 32, 30, 31, 30, 29, 30, 29, 30, 30},
}

var MonthNames = [12]string{
	"Baisakh", "Jestha", "Ashadh", "Shrawan", "Bhadra", "Ashwin", "Kartik",
	"Mangsir", "Paush", "Mangh", "Falgun", "Chaitra",
}

func New(Day int, Month interface{}, Year int) (Date, error) {
	var MonthInt int
	switch Month.(type) {
	case string:
		for i, v := range MonthNames {
			if v == Month.(string) {
				MonthInt = i + 1
				break
			}
		}
	case int:
		MonthInt = Month.(int)
	default:
		return nil, errors.New("month has to be of value int or string")
	}
	d := date{
		Day:   Day,
		Month: MonthInt,
		Year:  Year,
	}
	if !d.isValid() {
		return nil, errors.New("not a valid date")
	}
	return d, nil
}

func NewFromGregorian(gregorianDay, gregorianMonth, gregorianYear int) (Date, error) {
	var bsYear = gregorianYear + 56         //first rough calculation, might become 57 later
	var bsMonth = 9                         //Jan 1 always fall in BS month Paush which is the 9th month
	var daysSinceJanFirstToEndOfBsMonth int //days calculated from 1st Jan till the end of the actual BS month,
	                                        // we use this value to check if the gregorian Date is in the actual BS month

	year := time.Date(gregorianYear, time.Month(gregorianMonth), gregorianDay, 0, 0, 0, 0, time.UTC)
	var gregorianDayOfYear = year.YearDay()

	//get the BS day in Paush (month 9) of 1st January
	var dayOfFirstJanInPaush = calendardata[bsYear][0]

	//check how many days are left of Paush
	daysSinceJanFirstToEndOfBsMonth = calendardata[bsYear][bsMonth] - dayOfFirstJanInPaush + 1

	//If the gregorian day-of-year is smaller or equal to the sum of days between the 1st January and
	//the end of the actual BS month we found the correct nepali month.
	//Example:
	//The 4th February 2011 is the gregorianDayOfYear 35 (31 days of January + 4)
	//1st January 2011 is in the BS year 2067 and its the 17th day of Paush (9th month)
	//In 2067 Paush had 30days, This means (30-17+1=14) there are 14days between 1st January and end of Paush
	//(including 17th January)
	//The gregorianDayOfYear (35) is bigger than 14, so we check the next month
	//The next BS month (Mangh) has 29 days
	//29+14=43, this is bigger than gregorianDayOfYear(35) so, we found the correct nepali month
	for ; gregorianDayOfYear > daysSinceJanFirstToEndOfBsMonth; {
		bsMonth++
		if bsMonth > 12 {
			bsMonth = 1
			bsYear++
		}
		daysSinceJanFirstToEndOfBsMonth += calendardata[bsYear][bsMonth]
	}

	//the last step is to calculate the nepali day-of-month
	//We know the correct BS month, and we know the days since 1st Jan till the end of this month.
	//Subtracting the day-of-the-year of the gregorian calendar from the days since 1st Jan till the end of the correct
	//BS month will give us the amount of days between the searched day and the end of the BS month.
	//Subtracting that number from the amount of days in the BS month should bring us to the correct date.
	//to continue our example from before:
	//we calculated there are 43 days from 1st. January (17 Paush) till end of Mangh (29 days)
	//when we subtract from this 43days the day-of-year of the the gregorian date (35), we know how far the searched day is away
	//from the end of the nepali month. So we simply subtract this number from the amount of days in this month (30)
	var bsDay = calendardata[bsYear][bsMonth] - (daysSinceJanFirstToEndOfBsMonth - gregorianDayOfYear)

	return New(bsDay, bsMonth, bsYear)
}

func (d date) GetDay() int {
	return d.Day
}

func (d date) GetMonth() int {
	return d.Month
}

func (d date) GetYear() int {
	return d.Year
}

func (d date) GetMonthName() string {
	return MonthNames[d.Month-1]
}

func (d date) isValid() bool {
	//some rough testing
	if d.Day <= 0 || d.Day > 32 || d.Month <= 0 || d.Month > 12 || d.Year <= 0 {
		return false
	}
	//do we have data of that year?
	if _, ok := calendardata[d.Year]; !ok {
		return false
	}
	//does that particular month have so many days?
	if d.Day > calendardata[d.Year][d.Month] {
		return false
	}
	return true
}

func (d date) GetGregorianDate() (time.Time, error)  {
	var daysAfterJanFirstOfGregorianYear = 0 //we will add all the days that went by since the 1st.
	                                         //January and then we can get the gregorian Date
	var gregorianYear int
	var nepaliMonthToCheck  = d.Month
	var nepaliYearToCheck  = d.Year


	//get the correct year
	//after the month of Paush (9) or in Paush but after 1st Jan in that BS year we have to subtract 56 years, else 57
	if d.Month > 9 || (d.Month == 9 && d.Day >= calendardata[d.Year][0]) {
		gregorianYear = d.Year - 56
	} else {
		gregorianYear = d.Year - 57
	}

	//first we add the amount of days in the actual BS month as the day of year in the gregorian one
	//because at least this days are gone since the 1st. Jan.
	if d.Month != 9 {
		daysAfterJanFirstOfGregorianYear = d.Day
		nepaliMonthToCheck--
	}

	//now we loop through all nepali months and add the amount of days to daysAfterJanFirstOfGregorianYear
	//we do this till we reach Paush (9th month). 1st. January always falls in this month
	for ; nepaliMonthToCheck != 9 ;nepaliMonthToCheck-- {
		if nepaliMonthToCheck <= 0 {
			nepaliMonthToCheck = 12
			nepaliYearToCheck--
			//do we have data of that year?
			if _, ok := calendardata[nepaliYearToCheck]; !ok {
				return time.Time{}, errors.New("cannot convert date, missing data")
			}
		}
		daysAfterJanFirstOfGregorianYear += calendardata[nepaliYearToCheck][nepaliMonthToCheck]
	}

	//If the date that has to be converted is in Paush (month no. 9) we have to do some other calculation
	if d.Month == 9 {
		//add the days that are passed since the first day of Paush and substract the amount of days that lie between
		//1st. Jan and 1st Paush
		daysAfterJanFirstOfGregorianYear += d.Day - calendardata[nepaliYearToCheck][0]

		//for the first days of Paush we have now negative values
		//so we calculate daysAfterJanFirstOfGregorianYear for the previous year
		//last-day of the year (365 or 366) plus the negative daysAfterJanFirstOfGregorianYear value
		if daysAfterJanFirstOfGregorianYear < 0 {
			year := time.Date(gregorianYear, time.December, 31, 0, 0, 0, 0, time.UTC)
			daysAfterJanFirstOfGregorianYear = year.YearDay() + daysAfterJanFirstOfGregorianYear
		}
	} else {
		//add the days of Paush that are after 1st Jan
		daysAfterJanFirstOfGregorianYear += calendardata[nepaliYearToCheck][9] - calendardata[nepaliYearToCheck][0]
	}

	gregorianDate := time.Date(gregorianYear,1,1,0,0,0,0,time.UTC)
	gregorianDate = gregorianDate.AddDate(0,0, daysAfterJanFirstOfGregorianYear)
	return gregorianDate, nil
}
