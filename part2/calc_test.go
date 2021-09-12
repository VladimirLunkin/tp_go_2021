package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		expect    int
		assertion require.ComparisonAssertionFunc
	}{
		{"Два чсила", "1 + 2", 3, require.Equal},
		{"Три числа", "1 + 2 + 3", 6, require.Equal},
		{"Многозначные числа", "15000 + 200 - 30000000", -29984800, require.Equal},
		{"Лишние пробелы", "   1- 30 +  12    *2   ", -5, require.Equal},
		{`Выражение с префиксным "-"`, "-11111 * 2 - 20000", -42222, require.Equal},
		{"Скобки 1", "10 + (2 - 9) * 2", -4, require.Equal},
		{"Скобки 2", "(2 - 9) * (2 + 10)", -84, require.Equal},
		{"Скобки 3", "3 + 4 * 2 / (1 - 5)*2", -1, require.Equal},
		{"Вложенные скобки 1", "((10 + 2 * (2 - 9)) * 2)", -8, require.Equal},
		{"Вложенные скобки 2", "((1+2*3))*((2-1)/1)", 7, require.Equal},
		{`Скобки и знак минус 1`, "5 * (-2)", -10, require.Equal},
		{`Скобки и знак минус 2`, "-5 * (-2 + 12)", -50, require.Equal},
		{"Много действий", "(25 - 4) / 7 + 5 * (12 - 3) * 14/2 * (17 -8) + 3", 2841, require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := calc(tt.inputStr)
			tt.assertion(t, tt.expect, result)
		})
	}
}

func TestCalcWrongExpression(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		expect    int
		assertion require.ComparisonAssertionFunc
	}{
		{"Пустая строка", "", 0, require.Equal},
		{"Строка без чисел 1", "-", 0, require.Equal},
		{"Строка без чисел 2", "+", 0, require.Equal},
		{"Не валидные символы 1", "abc", 0, require.Equal},
		{"Не валидные символы 2", "1 + x", 0, require.Equal},
		{"Не валидные символы 3", "2 % 3", 0, require.Equal},
		{"Лишнии знаки 1", "1 + +2", 0, require.Equal},
		{"Лишнии знаки 2", "+7-13", 0, require.Equal},
		{"Лишнии знаки 3", "12 * +2", 0, require.Equal},
		{"Незакрытые скобки 1", "(1 +2", 0, require.Equal},
		{"Незакрытые скобки 2", "2 * (3*(1-14)", 0, require.Equal},
		{"Незакрытые скобки 3", "3 - 4+2)*6", 0, require.Equal},
		{"Незакрытые скобки 4", ")3 - 4+2)*6", 0, require.Equal},
		{"Незакрытые скобки 5", "3 - )4+2)*6", 0, require.Equal},
		{"Пустые скобки 1", "()", 0, require.Equal},
		{"Пустые скобки 2", "-()", 0, require.Equal},
		{"Унарный оператор", "-(2)", 0, require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := calc(tt.inputStr)
			tt.assertion(t, tt.expect, result)
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		expect    []string
		assertion require.ComparisonAssertionFunc
	}{
		{"", "1 + 2", []string{"1", "2", "+"}, require.Equal},
		{"", "1 + 2    + 3   ", []string{"1", "2", "+", "3", "+"}, require.Equal},
		{"", "15 + 200    + 3   ", []string{"15", "200", "+", "3", "+"}, require.Equal},
		{"", "3 + 4 * 2 / (1 - 5)*2", strings.Split("342*15-/2*+", ""), require.Equal},
		{"", "((1+2*3))*((2-1)/1)", strings.Split("123*+21-1/*", ""), require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := convertExpToRPN(tt.inputStr)
			tt.assertion(t, tt.expect, c)
		})
	}
}
