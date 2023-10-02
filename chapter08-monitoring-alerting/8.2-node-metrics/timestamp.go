package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TimeUnit represent time unit
type TimeUnit string

const (
	// TimeUnitMilliSeconds is milli seconds unit
	TimeUnitMilliSeconds TimeUnit = "milliseconds"
	// TimeUnitSeconds is the seconds unit
	TimeUnitSeconds TimeUnit = "seconds"
	// TimeUnitMinutes is the minutes unit
	TimeUnitMinutes TimeUnit = "minutes"
	// TimeUnitHours is the hours unit
	TimeUnitHours TimeUnit = "hours"
	// TimeUnitDays is the days unit
	TimeUnitDays TimeUnit = "days"
	// TimeUnitMonths is the months unit
	TimeUnitMonths TimeUnit = "months"
)

// Timestamp is the type for time.Time
type Timestamp time.Time

// Timestamp layouts
const (
	// IMPORTANT NOTE: When add new timestamp layout,
	// we must modify MIN_TIMESTAMP_LEN and MAX_TIMESTAMP_LEN according to new layout
	// when evaulate, using regex, we evaluate from top to bottom as declaring in this section,
	// from more generic to more specific according to declare in consts section
	MIN_TIMESTAMP_LEN           int    = 10 // LEN(2006-01-02)
	MAX_TIMESTAMP_LEN           int    = 33 // LEN(2006-01-02 15:04:05.000 +0000 UTC)
	TimestampLayout             string = "2006-01-02 15:04:05"
	TimestampLayoutMs           string = "2006-01-02 15:04:05.000"
	TimestampLayoutWithoutTime  string = "2006-01-02"
	TimestampLayoutNoMinute     string = "2006-01-02 15"
	TimestampLayoutNoSecond     string = "2006-01-02 15:04"
	TimestampLayoutISO8601      string = "2006-01-02T15:04:05Z"
	TimestampLayoutISO8601NoZ   string = "2006-01-02T15:04:05"
	TimestampLayoutISO8601MsNoZ string = "2006-01-02T15:04:05.0000000"
	TimestampLayoutWithUTC      string = "2006-01-02 15:04:05 +0000 UTC"
	TimestampLayoutMsWithUTC    string = "2006-01-02 15:04:05.000 +0000 UTC"
)

var (
	TimestampLayoutLen             = len(TimestampLayout)
	TimestampLayoutMsLen           = len(TimestampLayoutMs)
	TimestampLayoutWithoutTimeLen  = len(TimestampLayoutWithoutTime)
	TimestampLayoutNoMinuteLen     = len(TimestampLayoutNoMinute)
	TimestampLayoutNoSecondLen     = len(TimestampLayoutNoSecond)
	TimestampLayoutISO8601Len      = len(TimestampLayoutISO8601)
	TimestampLayoutISO8601NoZLen   = len(TimestampLayoutISO8601NoZ)
	TimestampLayoutISO8601MsNoZLen = len(TimestampLayoutISO8601MsNoZ)
	TimestampLayoutWithUTCLen      = len(TimestampLayoutWithUTC)
	TimestampLayoutMsWithUTCLen    = len(TimestampLayoutMsWithUTC)
)

// timestampLayout will evaluate layout in order
func timestampLayout(s string) string {
	sLen := len(s)

	// Optimize: for string which is not timestamp, no need to be evaluate with regex
	if sLen < MIN_TIMESTAMP_LEN {
		return ""
	}
	if sLen > MAX_TIMESTAMP_LEN {
		return ""
	}
	// Optimize: every timestamp format must start with YYYY-MM-DD 2006-01-02
	// there must be - at position 4 and 7
	if s[4] != '-' || s[7] != '-' {
		return ""
	}

	isMatch := false

	// We need to evaluate every length of timestamp layout because some of timestamp layout has same len
	// Anyway we evaulate from top to bottom, from more generic to more specific according to declare in consts section
	if sLen == TimestampLayoutLen {
		isMatch = getTimestampLayoutRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayout
	}

	if sLen == TimestampLayoutMsLen {
		isMatch = getTimestampLayoutMsRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutMs
	}

	if sLen == TimestampLayoutWithoutTimeLen {
		isMatch = getTimestampLayoutWithoutTimeRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutWithoutTime
	}

	if sLen == TimestampLayoutNoMinuteLen {
		isMatch = getTimestampLayoutNoMinuteRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutNoMinute
	}

	if sLen == TimestampLayoutNoSecondLen {
		isMatch = getTimestampLayoutNoSecondRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutNoSecond
	}

	if sLen == TimestampLayoutISO8601Len {
		isMatch = getTimestampLayoutISO8601Regex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutISO8601
	}

	if sLen == TimestampLayoutISO8601NoZLen {
		isMatch = getTimestampLayoutISO8601NoZRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutISO8601NoZ
	}

	if sLen == TimestampLayoutISO8601MsNoZLen {
		isMatch = getTimestampLayoutISO8601MsNoZRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutISO8601MsNoZ
	}

	if sLen == TimestampLayoutWithUTCLen {
		isMatch = getTimestampLayoutWithUTCRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutWithUTC
	}

	if sLen == TimestampLayoutMsWithUTCLen {
		isMatch = getTimestampLayoutMsWithUTCRegex().MatchString(s)
	}
	if isMatch {
		return TimestampLayoutMsWithUTC
	}

	return ""
}

// Errors
var (
	ErrTimeStringIsRequired = errors.New("time string is required")
)

type InvalidTimestampFormatError struct {
	t string
}

func NewInvalidTimestampFormatError(time string) *InvalidTimestampFormatError {
	return &InvalidTimestampFormatError{time}
}

func (e *InvalidTimestampFormatError) Error() string {
	return fmt.Sprintf("invalid timestamp format: %s", e.t)
}

// NewTimestampNow return timestamp of Now
func NewTimestampNow() *Timestamp {
	// When we run unit test (PAM4_ENV=test),
	// we can mock this function via env using name of function, eg, Mock_NewTimestampNow=2020-02-01 20:20:20
	env := os.Getenv("PAM4_ENV")
	if env == "test" {
		nowTime := os.Getenv("Mock_NewTimestampNow")
		if len(nowTime) > 0 {
			ts, _ := NewTimestamp(nowTime)
			return ts
		}
	}
	// For other env, this function return now in UTC
	tm := time.Now()
	ts := Timestamp(tm)
	return &ts
}

// NewTimestampNow return timestamp of Now
func NewTimestampToday() *Timestamp {
	// When we run unit test (PAM4_ENV=test),
	// we can mock this function via env using name of function, eg, Mock_NewTimestampNow=2020-02-01 20:20:20
	env := os.Getenv("PAM4_ENV")
	if env == "test" {
		todayTime := os.Getenv("Mock_NewTimestampToday")
		if len(todayTime) > 0 {
			ts, _ := NewTimestamp(todayTime)
			return ts
		}
	}
	// For other env, this function return today in UTC (time with Hour/Minute/Sec/Ms set to 0)
	tm := time.Now()
	td := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	ts := Timestamp(td)
	return &ts
}

// NewTimestampUTCNow return timestamp of utc now
func NewTimestampUTCNow() *Timestamp {
	// When we run unit test (PAM4_ENV=test),
	// we can mock this function via env using name of function, eg, Mock_NewTimestampUTCNow=2020-02-01 20:20:20
	env := os.Getenv("PAM4_ENV")
	if env == "test" {
		nowTime := os.Getenv("Mock_NewTimestampUTCNow")
		if len(nowTime) > 0 {
			ts, _ := NewTimestamp(nowTime)
			return ts
		}
	}

	location := time.FixedZone("UTC+0", 0)
	tm := time.Now().In(location)
	ts := Timestamp(tm)
	return &ts
}

func NewTimestampFromMs(ms string) (*Timestamp, error) {
	if len(ms) == 0 {
		return nil, ErrTimeStringIsRequired
	}
	millis, err := strconv.ParseInt(string(ms), 10, 64)
	if err != nil {
		return nil, &InvalidTimestampFormatError{ms}
	}
	t := time.Unix(0, millis*int64(time.Millisecond))
	return NewTimestampT(t), nil
}

// TimestampZero is the default value for timestamp
// time.Parse(timestampLayout, "0001-01-01 00:00:00")
// var TimestampZero Timestamp

// NewTimestamp from string
func NewTimestamp(t string) (*Timestamp, error) {
	if len(t) == 0 {
		return nil, ErrTimeStringIsRequired
	}
	t = strings.TrimSpace(t)
	layout := timestampLayout(t)
	if layout == "" {
		return nil, &InvalidTimestampFormatError{t}
	}
	tm, err := time.Parse(layout, t)
	if err != nil {
		return nil, err
	}
	ts := Timestamp(tm)
	return &ts, nil
}

// NewTimestampIn parse time string to timestamp in location
func NewTimestampIn(t string, timezone string, offset int) (*Timestamp, error) {
	if t == "" {
		return nil, ErrTimeStringIsRequired
	}
	layout := timestampLayout(t)
	if layout == "" {
		return nil, &InvalidTimestampFormatError{t}
	}
	location := time.FixedZone(timezone, offset*60*60) // FixedZone receive offset in seconds
	tm, err := time.ParseInLocation(layout, t, location)
	if err != nil {
		return nil, err
	}
	ts := Timestamp(tm)
	return &ts, nil
}

// NewTimestampT cast from time to Timestamp
func NewTimestampT(t time.Time) *Timestamp {
	tm := Timestamp(t)
	return &tm
}

func (t *Timestamp) IsSameDay(with *Timestamp) bool {
	if with == nil {
		return false
	}
	tm := t.Time()
	wtm := with.Time()

	return (tm.Year() == wtm.Year()) && (tm.Month() == wtm.Month()) && (tm.Day() == wtm.Day())
}

// Diff return positive number when diffWith less than t
// Diff return negative number when diffWith greater than t
// Return value will be in TimeUnit unit
// Return value will round up to integer number
// Default return value is in Nanoseconds unit
func (t *Timestamp) Diff(diffWith *Timestamp, unit TimeUnit) int64 {
	tm := t.Time()
	duration := tm.Sub(*diffWith.Time())

	if unit == TimeUnitSeconds {
		res := duration.Seconds()
		if res < 0 {
			return int64(math.Floor(res))
		}
		return int64(math.Ceil(res))
	} else if unit == TimeUnitMinutes {
		res := duration.Minutes()
		if res < 0 {
			return int64(math.Floor(res))
		}
		return int64(math.Ceil(res))
	} else if unit == TimeUnitHours {
		res := duration.Hours()
		if res < 0 {
			return int64(math.Floor(res))
		}
		return int64(math.Ceil(res))
	} else if unit == TimeUnitDays {
		res := duration.Hours() / 24
		if res < 0 {
			return int64(math.Floor(res))
		}
		return int64(math.Ceil(res))
	} else if unit == TimeUnitMonths {
		res := duration.Hours() / 24 / 30
		if res < 0 {
			return int64(math.Floor(res))
		}
		return int64(math.Ceil(res))
	}

	return duration.Nanoseconds() / int64(time.Millisecond)
}

// Time return time.Time
func (t *Timestamp) Time() *time.Time {
	tm := time.Time(*t)
	return &tm
}

// Add add duration and return new timestamp
func (t *Timestamp) Add(duration time.Duration) *Timestamp {
	return NewTimestampT(t.Time().Add(duration))
}

// Convert timm zone and return new timestamp
func (t *Timestamp) ToTimezone(offset int) *Timestamp {
	_, zoneOffset := t.Time().Zone()

	// Standard offset must be between -12 and +14
	if zoneOffset == 0 && offset >= -12 && offset <= 14 && offset != 0 {
		symbol := "+"
		if offset < 0 {
			symbol = "-"
		}

		location := time.FixedZone(fmt.Sprintf("UTC%s%d", symbol, offset), offset*60*60)
		tm := t.Time().In(location)
		newTS := Timestamp(tm)
		return &newTS
	}

	oldTS := *t
	return &oldTS
}

// UnmarshalJSON transform string to Timestamp (use via golang reflection when parsing JSON)
func (t *Timestamp) UnmarshalJSON(src []byte) error {
	str := strings.Trim(string(src), "\"")
	if len(string(str)) == 0 {
		return nil
	}

	ts, err := NewTimestamp(string(str))
	if err != nil {
		return err
	}
	*t = *ts
	return nil
}

// Do not removed
// String return string represent Timestamp (use via golang reflection in fmt package)
func (t *Timestamp) String() string {
	return t.DateTimeString()
}

// DateTimeString return date time in string format 2006-01-02 15:04:05
func (t *Timestamp) DateTimeString() string {
	return time.Time(*t).Format(TimestampLayout)
}

// DateTimeStringWithUTC return date time in string format 2006-01-02 15:04:05 +0000 UTC
func (t *Timestamp) DateTimeStringWithUTC() string {
	return time.Time(*t).Format(TimestampLayoutWithUTC)
}

// DateTimeStringRoundSecondToZero return date time in string format 2006-01-02 15:04:00, and round seconds part to zero
func (t *Timestamp) DateTimeStringRoundSecondToZero() string {
	tm := t.Time()
	// 2006-01-02 15:04:00
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute())
}

// DateTimeStringMs return date time in string format 2006-01-02 15:04:05.999
func (t *Timestamp) DateTimeStringMs() string {
	return time.Time(*t).Format(TimestampLayoutMs)
}

// DateString return date part of timestamp as string in format 2006-01-02
func (t *Timestamp) DateString() string {
	return time.Time(*t).Format(TimestampLayoutWithoutTime)
}

// ShortDateString return date part of timestamp as string in format MM-DD
func (t *Timestamp) ShortDateString() string {
	str := time.Time(*t).Format(TimestampLayoutWithoutTime)
	// Remove year part
	return str[5:]
}

// ShortDateTimeString return date part of timestamp as string in format MM-DD HH:mm
func (t *Timestamp) ShortDateTimeString() string {
	str := time.Time(*t).Format(TimestampLayoutNoSecond)
	// Remove year part
	return str[5:]
}

// HourString return HH:00
func (t *Timestamp) HourString() string {
	tm := t.Time()
	return fmt.Sprintf("%02d:00", tm.Hour())
}

// WeekInYearString of 2021-05-13 is 2021-22 (week in year)
func (t *Timestamp) WeekInYearString() string {
	y, m := time.Time(*t).ISOWeek()
	return fmt.Sprintf("%d-%d", y, m)
}

// WeekInMonthString of 2021-05-13 is May-2 (week in month)
func (t *Timestamp) WeekInMonthString() string {
	// _, w := time.Time(*t).ISOWeek()

	// monthSeq := w / 4
	// wSeq := w % 4
	// // If wSeq == 0, it means this week is the last week of the month.
	// // else it is the week of next month.
	// if wSeq == 0 {
	// 	wSeq = 4
	// } else {
	// 	monthSeq += 1
	// }
	// monthName := time.Month(monthSeq)

	// return fmt.Sprintf("%s-%d", monthName.String()[:3], wSeq)

	tm := t.Time()
	monthName := time.Month(tm.Month())

	wSeq := 0
	if tm.Day() >= 1 && tm.Day() <= 7 {
		wSeq = 1
	} else if tm.Day() >= 8 && tm.Day() <= 14 {
		wSeq = 2
	} else if tm.Day() >= 15 && tm.Day() <= 21 {
		wSeq = 3
	} else if tm.Day() >= 22 && tm.Day() <= 31 {
		wSeq = 4
	}

	return fmt.Sprintf("%s-%d", monthName.String()[:3], wSeq)
}

func (t *Timestamp) YearMonthString() string {
	tm := t.Time()
	return fmt.Sprintf("%04d-%02d", tm.Year(), tm.Month())
}

// YearMonthAlias return year and month in format 200601
func (t *Timestamp) YearMonthAlias() string {
	tm := t.Time()
	return fmt.Sprintf("%04d%02d", tm.Year(), tm.Month())
}

// YearMonthDayAlias return year, month and day in format 20060102
func (t *Timestamp) YearMonthDayAlias() string {
	tm := t.Time()
	return fmt.Sprintf("%04d%02d%02d", tm.Year(), tm.Month(), tm.Day())
}

// YearMonthTimeAlias return year, month and time in format 20060102150405
func (t *Timestamp) YearMonthTimeAlias() string {
	tm := t.Time()
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

// Milliseconds return current timestamp in milliseconds
func (t *Timestamp) Milliseconds() int64 {
	return t.Time().UnixNano() / int64(time.Millisecond)
}

// Unix timestamp is Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
// UnixSeconds return unix timestamp in seconds
func (t *Timestamp) UnixSeconds() int {
	return int(t.Time().Unix())
}

// Unix timestamp is Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
// UnixHours return unix timestamp in hours
func (t *Timestamp) UnixHours() int {
	return t.UnixSeconds() / (60 * 60)
}

// MarshalJSON tranform Timestamp to string (use via golang reflection when parsing JSON)
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`"%s"`, time.Time(*t).Format(TimestampLayout))
	return []byte(str), nil
}

// Value convert timestampe to sql value (use via golang reflection when send data to RDBMS via ORM)
func (t Timestamp) Value() (driver.Value, error) {
	ts := t.Time()
	return *ts, nil
}

// Scan db datetime to timestamp (use via golang reflection when query using ORM)
func (t *Timestamp) Scan(value interface{}) error {
	if value == nil {
		t = nil
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*t = *NewTimestampT(v)
	}
	return nil
}

var timestampLayoutRegex *regexp.Regexp

func getTimestampLayoutRegex() *regexp.Regexp {
	if timestampLayoutRegex == nil {
		// Validate format 2006-01-02 15:04:05
		timestampLayoutRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]$`)
	}
	return timestampLayoutRegex
}

var timestampLayoutMsRegex *regexp.Regexp

func getTimestampLayoutMsRegex() *regexp.Regexp {
	if timestampLayoutMsRegex == nil {
		// Validate format 2006-01-02 15:04:05.000
		timestampLayoutMsRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]\.[0-9][0-9][0-9]$`)
	}
	return timestampLayoutMsRegex
}

var timestampLayoutWithoutTimeRegex *regexp.Regexp

func getTimestampLayoutWithoutTimeRegex() *regexp.Regexp {
	if timestampLayoutWithoutTimeRegex == nil {
		// Validate format 2006-01-02
		timestampLayoutWithoutTimeRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`)
	}
	return timestampLayoutWithoutTimeRegex
}

var timestampLayoutNoMinuteRegex *regexp.Regexp

func getTimestampLayoutNoMinuteRegex() *regexp.Regexp {
	if timestampLayoutNoMinuteRegex == nil {
		// Validate format 2006-01-02 15
		timestampLayoutNoMinuteRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9])$`)
	}
	return timestampLayoutNoMinuteRegex
}

var timestampLayoutNoSecondRegex *regexp.Regexp

func getTimestampLayoutNoSecondRegex() *regexp.Regexp {
	if timestampLayoutNoSecondRegex == nil {
		// Validate format 2006-01-02 15:04
		timestampLayoutNoSecondRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]$`)
	}
	return timestampLayoutNoSecondRegex
}

var timestampLayoutISO8601Regex *regexp.Regexp

func getTimestampLayoutISO8601Regex() *regexp.Regexp {
	if timestampLayoutISO8601Regex == nil {
		// Validate format 2006-01-02T15:04:05Z
		timestampLayoutISO8601Regex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T(2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]Z$`)
	}
	return timestampLayoutISO8601Regex
}

var timestampLayoutISO8601NoZRegex *regexp.Regexp

func getTimestampLayoutISO8601NoZRegex() *regexp.Regexp {
	if timestampLayoutISO8601NoZRegex == nil {
		// Validate format 2006-01-02T15:04:05
		timestampLayoutISO8601NoZRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T(2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]$`)
	}
	return timestampLayoutISO8601NoZRegex
}

var timestampLayoutISO8601MsNoZRegex *regexp.Regexp

func getTimestampLayoutISO8601MsNoZRegex() *regexp.Regexp {
	if timestampLayoutISO8601MsNoZRegex == nil {
		// Validate format 2006-01-02T15:04:05.0000000
		timestampLayoutISO8601MsNoZRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T(2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]\.[0-9][0-9][0-9][0-9][0-9][0-9][0-9]$`)
	}
	return timestampLayoutISO8601MsNoZRegex
}

var timestampLayoutWithUTCRegex *regexp.Regexp

func getTimestampLayoutWithUTCRegex() *regexp.Regexp {
	if timestampLayoutWithUTCRegex == nil {
		// Validate format 2006-01-02 15:04:05 +0000 UTC
		timestampLayoutWithUTCRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9] \+0000 UTC$`)
	}
	return timestampLayoutWithUTCRegex
}

var timestampLayoutMsWithUTCRegex *regexp.Regexp

func getTimestampLayoutMsWithUTCRegex() *regexp.Regexp {
	if timestampLayoutMsWithUTCRegex == nil {
		// Validate format 2006-01-02 15:04:05.000 +0000 UTC
		timestampLayoutMsWithUTCRegex = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]\.[0-9][0-9][0-9] \+0000 UTC$`)
	}
	return timestampLayoutMsWithUTCRegex
}
