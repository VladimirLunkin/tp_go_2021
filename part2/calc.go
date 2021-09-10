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

func parseStr(expressionStr string) (expression []string, err error) {
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
			} else if symbol == "-" {
				number = "-"
			} else {
				expression = append(expression, symbol)
			}
		} else {
			return nil, fmt.Errorf("invalid character: %s", symbol)
		}

	}
	if number != "" {
		expression = append(expression, number)
	}

	return
}

func convertExpToRPN(expressionStr string) (rpn []string, err error) {
	priorityOfOperations := map[string]uint{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 3,
		"*": 4,
		"/": 4,
	}

	var stk stack.Stack

	expression, err := parseStr(expressionStr)
	if err != nil {
		return nil, fmt.Errorf("parsing an expression: %s", err)
	}

	for _, currLexeme := range expression {
		if isNumber(currLexeme) {
			rpn = append(rpn, currLexeme)
		} else if currLexeme == "(" {
			stk.Push(currLexeme)
		} else if currLexeme == ")" {
			if stk.IsEmpty() {
				return nil, fmt.Errorf("missing opening parenthesis")
			}

			topStack, err := stk.Pop()
			if err != nil {
				return nil, err
			}
			for topStack != "(" {
				rpn = append(rpn, topStack)
				topStack, err = stk.Pop()
				//topStack, err = appendTopStack(rpn, stk)
				if err != nil {
					return nil, fmt.Errorf("missing opening parenthesis: %s", err)
				}
			}
		} else {
			if !stk.IsEmpty() {
				topStack, err := stk.Top()
				if err != nil {
					return nil, err
				}
				if priorityOfOperations[topStack] >= priorityOfOperations[currLexeme] {
					rpn = append(rpn, topStack)
					_, err = stk.Pop()
					if err != nil {
						return nil, err
					}
				}
			}
			stk.Push(currLexeme)
		}
	}

	for !stk.IsEmpty() {
		topStack, err := stk.Pop()
		if err != nil {
			return nil, err
		}
		rpn = append(rpn, topStack)
	}

	return
}

func mathOperationsOnStrings(leftArgumentStr, sign, rightArgumentStr string) (resultStr string, err error) {
	leftArgument, err := strconv.Atoi(leftArgumentStr)
	if err != nil {
		return "", err
	}
	rightArgument, err := strconv.Atoi(rightArgumentStr)
	if err != nil {
		return "", err
	}

	var result int
	switch sign {
	case "+":
		result = leftArgument + rightArgument
	case "-":
		result = leftArgument - rightArgument
	case "*":
		result = leftArgument * rightArgument
	case "/":
		result = leftArgument / rightArgument
	}

	return strconv.Itoa(result), nil
}

func calcRPN(rpn []string) (int, error) {
	var stk stack.Stack

	for _, currValue := range rpn {
		if isNumber(currValue) {
			stk.Push(currValue)
		} else {
			r, _ := stk.Pop() // TODO
			l, _ := stk.Pop() // TODO
			R, err := mathOperationsOnStrings(l, currValue, r)
			if err != nil {
				return 0, err
			}
			stk.Push(R)
		}
	}

	result, err := stk.Pop()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

func calc(expression string) (result int, err error) {
	rpn, err := convertExpToRPN(expression)
	if err != nil {
		return 0, fmt.Errorf("error converting expression to rpn: %s", err)
	}

	result, err = calcRPN(rpn)
	if err != nil {
		return 0, fmt.Errorf("rpn computation error: %s", err)
	}

	return
}

func readLineStdin() (line string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return
}

func main() {
	expression, err := readLineStdin()
	if err != nil {
		log.Fatal(err)
	}

	result, err := calc(expression)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

// TODO
// 2. Добавить функционал ошибок в функции calcRPN,
// 3. Тесты с ошибками
// 4. Рефакторинг
// 5. Чек вк
// 6. Пуш
