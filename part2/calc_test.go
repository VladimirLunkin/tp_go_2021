package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type CheckStructInput struct {
	name    string
	expression   string
}
type CheckStructOutput struct {
	expected int
}

type testCaseStruct struct {
	InputData  CheckStructInput
	OutputData CheckStructOutput
	assertion  require.ComparisonAssertionFunc
}

func TestCalc(t *testing.T) {
	tests := []testCaseStruct{
		{
			InputData: CheckStructInput{
				name: "Два чсила",
				expression: "1 + 2",
			},
			OutputData: CheckStructOutput{
				expected: 3,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Три чсила",
				expression: "1 + 2 + 3",
			},
			OutputData: CheckStructOutput{
				expected: 6,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Многозначные числа",
				expression: "15000 + 200 - 30000000",
			},
			OutputData: CheckStructOutput{
				expected: -29984800,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Лишние пробелы",
				expression: "   1- 30 +  12    *2   ",
			},
			OutputData: CheckStructOutput{
				expected: -5,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Скобки 1",
				expression: "10 + (2 - 9) * 2",
			},
			OutputData: CheckStructOutput{
				expected: -4,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Скобки 2",
				expression: "(2 - 9) * (2 + 10)",
			},
			OutputData: CheckStructOutput{
				expected: -84,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Скобки 3",
				expression: "3 + 4 * 2 / (1 - 5)*2",
			},
			OutputData: CheckStructOutput{
				expected: -1,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Вложенные скобки 1",
				expression: "((10 + 2 * (2 - 9)) * 2)",
			},
			OutputData: CheckStructOutput{
				expected: -8,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Вложенные скобки 2",
				expression: "((1+2*3))*((2-1)/1)",
			},
			OutputData: CheckStructOutput{
				expected: 7,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Много действий",
				expression: "(25 - 4) / 7 + 5 * (12 - 3) * 14/2 * (17 -8) + 3",
			},
			OutputData: CheckStructOutput{
				expected: 2841,
			},
			assertion: require.Equal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.InputData.name, func(t *testing.T) {
			result, _ := calc(tt.InputData.expression)
			tt.assertion(t, tt.OutputData.expected, result)
		})
	}
}

func TestCalcWrongExpression(t *testing.T) {
	tests := []testCaseStruct{
		{
			InputData: CheckStructInput{
				name: "Пустая строка",
				expression: "",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Строка без чисел 1",
				expression: "-",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Строка без чисел 2",
				expression: "+",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Не валидные символы 1",
				expression: "abc",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Не валидные символы 2",
				expression: "1 + x",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Не валидные символы 3",
				expression: "2 % 3",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Лишний знак 1",
				expression: "1 + -2",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Лишний знак 2",
				expression: "+7 - 13",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Лишний знак 3",
				expression: "12 * -2",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Незакрытые скобки 1",
				expression: "(1 +2",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Незакрытые скобки 2",
				expression: "2 * (3*(1-14)",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Незакрытые скобки 3",
				expression: "3 - 4+2)*6",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Незакрытые скобки 4",
				expression: ")3 - 4+2)*6",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Незакрытые скобки 5",
				expression: "3 - )4+2)*6",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Пустые скобки 1",
				expression: "()",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name: "Пустые скобки 2",
				expression: "-()",
			},
			OutputData: CheckStructOutput{
				expected: 0,
			},
			assertion: require.Equal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.InputData.name, func(t *testing.T) {
			result, _ := calc(tt.InputData.expression)
			tt.assertion(t, tt.OutputData.expected, result)
		})
	}
}
