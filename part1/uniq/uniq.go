package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	//arg1
	i         bool
	f         bool
	numFields int
	s         bool
	numChars  int
}

type repeatingLine struct {
	line                string
	numberOfRepetitions uint
}

func countDuplicateLines(inputRowset []string, isEqual func(l, r string, opt Options) bool, opt Options) []repeatingLine {
	var noDuplicateLines []repeatingLine

	i := -1
	for _, currLine := range inputRowset {
		if len(noDuplicateLines) != 0 && isEqual(noDuplicateLines[i].line, currLine, opt) {
			noDuplicateLines[i].numberOfRepetitions++
		} else {
			noDuplicateLines = append(noDuplicateLines, repeatingLine{currLine, 1})
			i++
		}
	}

	return noDuplicateLines
}

func applyKeyI(inputString string) string {
	return strings.ToLower(inputString)
}
func applyKeyF(inputString string, numFields int) string {
	return strings.Join(strings.Split(inputString, " ")[numFields:], " ")
}
func applyKeyS(inputString string, numChars int) string {
	return inputString[numChars:]
}

func isEqual(l, r string, opt Options) bool {
	if opt.i {
		l, r = applyKeyI(l), applyKeyI(r)
	}

	if opt.f {
		l, r = applyKeyF(l, opt.numFields), applyKeyF(r, opt.numFields)
	}

	if opt.s {
		l, r = applyKeyS(l, opt.numChars), applyKeyS(r, opt.numChars)
	}

	return l == r
}

func Uniq(inputRowset []string, opt Options) (string, error) {
	if len(inputRowset) == 0 || inputRowset[0] == "" {
		return "", nil
	}

	res := countDuplicateLines(inputRowset, isEqual, opt)

	// TODO сформировать вывод исходя из опций
	var outputStr string

	for _, currStr := range res {
		outputStr += strconv.Itoa(int(currStr.numberOfRepetitions)) + " " + currStr.line + "\n"
	}

	return outputStr, nil
}

// TODO разобраться с опциями, флагами добавить ввод/вывод
