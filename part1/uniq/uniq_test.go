package uniq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type CheckStructInput struct {
	name    string
	lines   []string
	options Options
}
type CheckStructOutput struct {
	expected string
}

type testCaseStruct struct {
	InputData  CheckStructInput
	OutputData CheckStructOutput
	assertion  require.ComparisonAssertionFunc
}

func TestUniq(t *testing.T) {
	tests := []testCaseStruct{
		{
			InputData: CheckStructInput{
				name:    "Без параметров",
				lines:   strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"),
				options: Options{},
			},
			OutputData: CheckStructOutput{
				expected: "I love music.\n\nI love music of Kartik.\nThanks.\nI love music of Kartik.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -c",
				lines:   strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"),
				options: Options{C: true},
			},
			OutputData: CheckStructOutput{
				expected: "3 I love music.\n1 \n2 I love music of Kartik.\n1 Thanks.\n2 I love music of Kartik.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -d",
				lines:   strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"),
				options: Options{D: true},
			},
			OutputData: CheckStructOutput{
				expected: "I love music.\nI love music of Kartik.\nI love music of Kartik.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -u",
				lines:   strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"),
				options: Options{U: true},
			},
			OutputData: CheckStructOutput{
				expected: "\nThanks.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -i",
				lines:   strings.Split("I LOVE MUSIC.\nI love music.\nI LoVe MuSiC.\n\nI love MuSIC of Kartik.\nI love music of kartik.\nThanks.\nI love music of kartik.\nI love MuSIC of Kartik.", "\n"),
				options: Options{I: true},
			},
			OutputData: CheckStructOutput{
				expected: "I LOVE MUSIC.\n\nI love MuSIC of Kartik.\nThanks.\nI love music of kartik.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -f num",
				lines:   strings.Split("We love music.\nI love music.\nThey love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.", "\n"),
				options: Options{NumFields: 1},
			},
			OutputData: CheckStructOutput{
				expected: "We love music.\n\nI love music of Kartik.\nThanks.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "С параметром -s num",
				lines:   strings.Split("I love music.\nA love music.\nC love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.", "\n"),
				options: Options{NumChars: 1},
			},
			OutputData: CheckStructOutput{
				expected: "I love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "Нет входных данных",
				lines:   []string{},
				options: Options{},
			},
			OutputData: CheckStructOutput{
				expected: "",
			},
			assertion: require.Equal,
		},
		{
			InputData: CheckStructInput{
				name:    "Неправильные параметры",
				lines:   strings.Split("I love music.\nI love music.", "\n"),
				options: Options{C: true, D: true, U: true},
			},
			OutputData: CheckStructOutput{
				expected: "",
			},
			assertion: require.Equal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.InputData.name, func(t *testing.T) {
			tt.assertion(t, tt.OutputData.expected, Uniq(tt.InputData.lines, tt.InputData.options))
		})
	}
}
