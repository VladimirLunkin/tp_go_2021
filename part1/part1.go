package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/VladimirLunkin/tp_go_2021/part1/uniq"
)

var options uniq.Options

func init() {
	const (
		flagC = "подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом."
		flagD = "вывести только те строки, которые повторились во входных данных."
		flagU = "вывести только те строки, которые не повторились во входных данных."
		flagI = "не учитывать регистр букв"
		flagF = "не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом."
		flagS = "не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля)."
	)
	flag.BoolVar(&options.C, "c", false, flagC)
	flag.BoolVar(&options.D, "d", false, flagD)
	flag.BoolVar(&options.U, "u", false, flagU)
	flag.BoolVar(&options.I, "i", false, flagI)
	flag.IntVar(&options.NumFields, "f", 0, flagF)
	flag.IntVar(&options.NumChars, "s", 0, flagS)
}

const ValidOptions = "options: [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

func readData() []string {
	inputFile := os.Stdin

	if flag.NArg() > 0 {
		var err error
		inputFile, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(inputFile)

	buf, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(buf), "\n")
}

func printResult(result string) {
	outputFile := os.Stdout

	if flag.NArg() == 2 {
		var err error
		outputFile, err = os.Create(flag.Arg(1))
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(outputFile)

	io.StringWriter.WriteString(outputFile, result)
}

func main() {
	flag.Parse()

	if !options.IsValid() || flag.NArg() > 2 {
		fmt.Println(ValidOptions)
		return
	}

	inputData := readData()

	result := uniq.Uniq(inputData, options)

	printResult(result)
}
