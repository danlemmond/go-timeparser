package timeparser

import (
    "fmt"
    "time"
	"regexp"
)

type parseTimeStruct struct {
	clockwords	[]string
	numbers		map[string]int
	unit		map[string]int
	countable	map[string]int
	month		map[string]string
	daytime		map[string]string
	fillwords	map[string]byte
}

var master = parseTimeStruct{clockwords, numbers, unit, countable, month, daytime, fillwords}

var parens = regexp.MustCompile(`["'<>\(\)]`)
var chars = regexp.MustCompile(`/(\d)([A-Za-z])/`)
var replace = regexp.MustCompile(`/([^\x00-\x7F])`)


var clockwords = []string{"oclock", "o\\clock"}

var numbers = map[string]int {
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
	"ten": 10,
	"eleven": 11,
	"twelve": 12,
	"thirteen": 13,
	"fourteen": 14,
	"fifteen": 15,
	"sixteen":16,
	"seventeen": 17,
	"eighteen": 18,
	"nineteen": 19,
	"twenty": 20,
	"thirty": 30,
	"forty": 40,
	"fifty": 50,
	"sixty": 60,
	"seventy": 70,
	"eighty": 80,
	"ninety": 90,
	"hundred": 100,
	"thousand": 1000,
	"million": 1000000,
}

var unit = map[string]int {
	"millisecond" : 1,
	"second" : 1000,
	"minute" : 60000,
	"hour" : 3600000,
	"day" : 86400000,
	"week" : 604800000,
	"month" : 2592000000,
	"quarter" : 7776000000,
	"year" : 31536000000,
	"decade": 315360000000,
}

var countable = map[string]int {
	"before yesterday" : -17280000,
	"yesterday" : -8640000,
	"today" : 1,
	"day after tomorrow" : 17280000,
	"tomorrow" : 8640000,
	"in a week" : 60480000,
	"last week" : -60480000,
}

var month = map[string]string {
	"jan" : "01",
	"feb" : "02",
	"mar" : "03",
	"apr" : "04",
	"may" : "05",
	"jun" : "06",
	"jul" : "07",
	"aug" : "08",
	"sep" : "09",
	"oct" : "10",
	"nov" : "11",
	"dec" : "12",
}

var daytime = map[string]string{
	"dawn": "04:00",
	"morning": "06:00",
	"afternoon": "15:00",
	"noon": "12:00",
	"midday": "12:00",
	"pre-evening": "17:00",
	"preevening": "17:00",
	"evening": "19:00",
	"dusk": "20:00",
	"midnight": "24:00",
	"night": "22:00",
}

var fillwords = map[string]byte {
	"ago": '-',
	"in": '+',
}

func parseTime(s string, n string) {
	var re, lang, encoded, timedif, unit, word, hhmmss, hhmmss2, tzoffset, timeregex, wordForNow, curLang string
	var integer, pbint, val int

	now := time.Now().Unix()

}

func adWordsToRegex(fillfoo byte, first int) {

}

