package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Id = int64

// Basic time defined in milliseconds.
type Time = int64

// We consider standard time.Time as a Date, other mentions of Time are ms
// timestamps.
type Date = time.Time
type Dict = map[string]any

func Unwrap(err error) {
	if err != nil {
		panic(err)
	}
}

func Print(obj ...any) {
	fmt.Println(obj...)
}

// Reference: https://stackoverflow.com/a/13295158/14748231
func TimeToDate(t Time) (Date, error) {
	return time.Unix(0, t*int64(time.Millisecond)), nil
}

func TimeNow() Time {
	return DateNow().UnixMilli()
}

func DateNow() Date {
	return time.Now()
}

type Code = int16
type Error struct {
	code Code
}

func (e *Error) Error() string {
	return fmt.Sprintf("[Error %d] %s", e.code, TranslateCode(e.code))
}

func (e *Error) Code() Code {
	return e.code
}

func NewError(code Code) *Error {
	return &Error{code}
}

type Locale = string
type TranslationKey = string

var translationMap = map[Locale]map[TranslationKey]string{}
var translationLocale string = "en"

const CodeOk Code = 0
const CodeError Code = 1

// Register a translation from a CSV file.
// CSV file structure:
// key(string),text(string)
//
// This function can be called many times, each new call the old matching
// entries will be overwritten.
//
// Text may contain placeholders in form of `%` to accept incoming value,
// which will always be converted to string.
//
// For list of locales refer to https://docs.godotengine.org/en/4.3/tutorials/i18n/locales.html
func RegisterTranslationCsv(path string, locale Locale, delimiter rune) error {
	file, e := os.Open(path)
	if e != nil {
		return e
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = delimiter
	records, e := reader.ReadAll()
	if e != nil {
		return e
	}

	localeMap, ok := translationMap[locale]
	if !ok {
		localeMap = map[TranslationKey]string{}
		translationMap[locale] = localeMap
	}

	for i, record := range records {
		if len(record) != 2 {
			return NewError(CodeError)
		}
		if i == 0 {
			continue
		}
		localeMap[record[0]] = record[1]
	}

	return nil
}

// Codes are translated using keys `CODE_%`, where `%` is the number.
func TranslateCode(code Code) string {
	return Translate(fmt.Sprintf(
		"CODE_%", strconv.Itoa(int(code)),
	))
}

func Translate(key TranslationKey, args ...any) string {
	localeMap, ok := translationMap[translationLocale]
	if !ok {
		return "???"
	}
	text, ok := localeMap[strings.ToUpper(key)]
	if !ok {
		return "???"
	}
	return text
}

// Logging implementation.
func Log(obj ...any) {
	// TODO: write to sink, which may be stderr
	fmt.Println(obj...)
}