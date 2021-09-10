package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/VladimirLunkin/tp_go_2021/part2/stack"
)

func isNumber(number string) bool {
	rNumber := regexp.MustCompile(`[\d]`)
	return rNumber.MatchString(number)
}
func isSign(sign string) bool {
	rSign := regexp.MustCompile(`[+\-*/()]`)
	return rSign.MatchString(sign)
}

func parseStr(expressionStr string) (expression []string) {
	rSpace := regexp.MustCompile(`\s+`)
	inputStrWithoutSpace := rSpace.ReplaceAllString(expressionStr, "")

	var number string
	for _, symbol := range strings.Split(inputStrWithoutSpace, "") {
		if isNumber(symbol) {
			number += symbol
		} else if isSign(symbol) {
			if isNumber(number) {
				expression = append(expression, number, symbol)
				number = ""
			} else if symbol == "-"{
				number = "-"
			} else {
				expression = append(expression, symbol)
			}
		} else {
			// TODO ложный символ
		}

	}
	if number != "" {
		expression = append(expression, number)
	}

	return
}

func convertExpToRPN(expressionStr string) (rpn []string) {
	priorityOfOperations := map[string]uint{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 3,
		"*": 4,
		"/": 4,
	}

	var stk stack.Stack
	rNumber := regexp.MustCompile(`[\d]`)
	for _, currLexeme := range parseStr(expressionStr) {
		if rNumber.MatchString(currLexeme) {
			rpn = append(rpn, currLexeme)
		} else if currLexeme == "(" {
			stk.Push(currLexeme)
		} else if currLexeme == ")" {
			if stk.IsEmpty() {
				// TODO вернуть ошибку
			}

			topStack := stk.Pop()
			for topStack != "(" && !stk.IsEmpty() {
				topStack, rpn = stk.Pop(), append(rpn, topStack)
			}

			if stk.IsEmpty() && topStack != "(" {
				// TODO вернуть ошибку
			}
		} else {
			if !stk.IsEmpty() {
				if priorityOfOperations[stk.Top()] >= priorityOfOperations[currLexeme] {
					rpn = append(rpn, stk.Pop())
				}
			}
			stk.Push(currLexeme)
		}
	}

	for !stk.IsEmpty() {
		rpn = append(rpn, stk.Pop())
	}

	return
}

func calcBinOperation(arg1Str, sign, arg2Str string) string {
	arg1, err := strconv.Atoi(arg1Str)
	if err != nil {
		// TODO error
		log.Fatal(err)
	}
	arg2, err := strconv.Atoi(arg2Str)
	if err != nil {
		// TODO error
		log.Fatal(err)
	}

	var result int
	switch sign {
	case "+":
		result = arg1 + arg2
	case "-":
		result = arg1 - arg2
	case "*":
		result = arg1 * arg2
	case "/":
		result = arg1 / arg2
	}

	return strconv.Itoa(result)
}

func calcRPN(rpn []string) (result int) {
	var stk stack.Stack

	for _, currValue := range rpn {
		if isNumber(currValue) {
			stk.Push(currValue)
		} else {
			r := stk.Pop()
			l := stk.Pop()
			stk.Push(calcBinOperation(l, currValue, r))
		}
	}

	result, _ = strconv.Atoi(stk.Pop())

	return result
}

func calc(expression string) int {
	rpn := convertExpToRPN(expression)

	return calcRPN(rpn)
}

func readLineStdin() (line string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	return
}

func main() {
	expression := readLineStdin()

	fmt.Println(calc(expression))
}

// TODO
// 2. Добавить функционал ошибок
// 3. Тесты с ошибками
// 4. Рефакторинг
// 5. Чек вк
// 6. Пуш
