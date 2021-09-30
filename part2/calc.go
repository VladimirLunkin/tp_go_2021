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

var (
	rNumber = regexp.MustCompile(`[\d]`)
	rSign = regexp.MustCompile(`[+\-*/]`)
)

func isNumber(number string) bool {
	return rNumber.MatchString(number)
}
func isSign(sign string) bool {
	return rSign.MatchString(sign)
}

// Из массива символов, начиная с позиции startPos, составляем число пока
// нам встречаются цифры. Возвращаем само число и позицию endPos, на которой оно закончилось.
func getNumberFromString(charArray []string, startPos int) (number string, endPos int) {
	for endPos = startPos; endPos < len(charArray); endPos++ {
		if isNumber(charArray[endPos]) {
			number += charArray[endPos]
		} else {
			endPos--
			return
		}
	}

	return
}

// Переводим входную строку в обратную польскую запись.
// В случае невалидного ввода, возвращаем ошибки
func convertExpToRPN(expressionStr string) (rpn []string, err error) {
	rSpace := regexp.MustCompile(`\s+`)
	expressionWithoutSpace := rSpace.ReplaceAllString(expressionStr, "")

	expression := strings.Split(expressionWithoutSpace, "")

	priorityOfOperations := map[string]uint{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 3,
		"*": 4,
		"/": 4,
	}
	var stk stack.Stack
	sign := ""

	for i := 0; i < len(expression); i++ {
		if expression[i] == "-" && (i == 0 || expression[i-1] == "(") {
			sign = "-"
		} else if isNumber(expression[i]) {
			number, endPos := getNumberFromString(expression, i)
			i = endPos

			number = sign + number
			sign = ""

			rpn = append(rpn, number)
		} else if expression[i] == "(" {
			stk.Push(expression[i])
		} else if expression[i] == ")" {
			topStack, err := stk.Pop()
			if err != nil {
				return nil, fmt.Errorf("missing opening parenthesis: %s", err)
			}
			for topStack != "(" {
				rpn = append(rpn, topStack)
				topStack, err = stk.Pop()
				if err != nil {
					return nil, fmt.Errorf("missing opening parenthesis: %s", err)
				}
			}
		} else if isSign(expression[i]) {
			if !stk.IsEmpty() {
				topStack, err := stk.Top()
				if err != nil {
					return nil, err
				}
				if priorityOfOperations[topStack] >= priorityOfOperations[expression[i]] {
					rpn = append(rpn, topStack)
					_, err = stk.Pop()
					if err != nil {
						return nil, err
					}
				}
			}
			stk.Push(expression[i])
		} else {
			return nil, fmt.Errorf("invalid character: %s", expression[i])
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

// Выполняет математическую операцию над аргументами типа String,
// возвращая результат вычисления тип String
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

// Получает на вход обратную польскую запись и вычисляет ее
func calcRPN(rpn []string) (int, error) {
	var stk stack.Stack

	for _, currValue := range rpn {
		if isNumber(currValue) {
			stk.Push(currValue)
		} else {
			rightArg, err := stk.Pop()
			if err != nil {
				return 0, err
			}
			leftArg, err := stk.Pop()
			if err != nil {
				return 0, err
			}

			resultStr, err := mathOperationsOnStrings(leftArg, currValue, rightArg)
			if err != nil {
				return 0, err
			}
			stk.Push(resultStr)
		}
	}

	resultStr, err := stk.Pop()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(resultStr)
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
