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

func (r Options) IsValid() bool {
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

func applyAllKeys(inputString string, options Options) (modifiedString string) {
	modifiedString = inputString
	if options.I {
		modifiedString = applyKeyI(modifiedString)
	}

	if options.NumFields != 0 {
		modifiedString = applyKeyF(modifiedString, options.NumFields)
	}

	if options.NumChars != 0 {
		modifiedString = applyKeyS(modifiedString, options.NumChars)
	}

	return modifiedString
}

type repeatingLine struct {
	line                string
	numberOfRepetitions uint
}

// Получает на вход набор строк inputRowset и опции options по которым сравниваются строки.
// На выходе массив из структур repeatingLine, каждая содержит строку и количество ее повторений в исходном наборе.
func countDuplicateLines(inputRowset []string, options Options) []repeatingLine {
	var repeatingLines []repeatingLine

	i := -1
	for _, currLine := range inputRowset {
		if len(repeatingLines) != 0 &&
			applyAllKeys(repeatingLines[i].line, options) == applyAllKeys(currLine, options) {
			repeatingLines[i].numberOfRepetitions++
		} else {
			repeatingLines = append(repeatingLines, repeatingLine{currLine, 1})
			i++
		}
	}

	return repeatingLines
}

func Uniq(inputRowset []string, options Options) string {
	if len(inputRowset) == 0 || !options.IsValid() {
		return ""
	}

	repeatingLines := countDuplicateLines(inputRowset, options)

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
