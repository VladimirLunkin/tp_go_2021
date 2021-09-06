package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	C, D, U   bool
	I         bool
	NumFields int
	NumChars  int
}

func (r Options) isValid() bool {
	return !(r.C && (r.D != r.U) || r.D && r.U) && r.NumFields >= 0 && r.NumChars >= 0
}

func applyKeyI(inputString string) string {
	return strings.ToLower(inputString)
}
func applyKeyF(inputString string, numFields int) string {
	tempStr := strings.Split(inputString, " ")
	if len(tempStr) <= numFields {
		return ""
	}
	return strings.Join(tempStr[numFields:], " ")
}
func applyKeyS(inputString string, numChars int) string {
	if len(inputString) <= numChars {
		return ""
	}
	return inputString[numChars:]
}

func isEqual(l, r string, options Options) bool {
	if options.I {
		l, r = applyKeyI(l), applyKeyI(r)
	}

	if options.NumFields != 0 {
		l, r = applyKeyF(l, options.NumFields), applyKeyF(r, options.NumFields)
	}

	if options.NumChars != 0 {
		l, r = applyKeyS(l, options.NumChars), applyKeyS(r, options.NumChars)
	}

	return l == r
}

type funcEqual func(l, r string, options Options) bool
type repeatingLine struct {
	line                string
	numberOfRepetitions uint
}

func countDuplicateLines(inputRowset []string, isEqual funcEqual, options Options) []repeatingLine {
	var repeatingLines []repeatingLine

	i := -1
	for _, currLine := range inputRowset {
		if len(repeatingLines) != 0 && isEqual(repeatingLines[i].line, currLine, options) {
			repeatingLines[i].numberOfRepetitions++
		} else {
			repeatingLines = append(repeatingLines, repeatingLine{currLine, 1})
			i++
		}
	}

	return repeatingLines
}

func Uniq(inputRowset []string, options Options) string {
	if len(inputRowset) == 0 || !options.isValid() {
		return ""
	}

	repeatingLines := countDuplicateLines(inputRowset, isEqual, options)

	var outputData string
	for _, currLines := range repeatingLines {
		if options.C {
			outputData += strconv.Itoa(int(currLines.numberOfRepetitions)) + " " + currLines.line + "\n"
		} else if !(options.D || options.U) ||
			(options.D && currLines.numberOfRepetitions > 1) ||
			(options.U && currLines.numberOfRepetitions == 1) {
			outputData += currLines.line + "\n"
		}
	}

	return outputData[:len(outputData)-1]
}
