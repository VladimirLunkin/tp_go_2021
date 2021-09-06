package uniq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name         string
		inputLines   []string
		inputOptions Options
		expect       string
		assertion    require.ComparisonAssertionFunc
	}{
		{"Без параметров", strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"), Options{}, "I love music.\n\nI love music of Kartik.\nThanks.\nI love music of Kartik.", require.Equal},
		{"С параметром -c", strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"), Options{C: true}, "3 I love music.\n1 \n2 I love music of Kartik.\n1 Thanks.\n2 I love music of Kartik.", require.Equal},
		{"С параметром -d", strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"), Options{D: true}, "I love music.\nI love music of Kartik.\nI love music of Kartik.", require.Equal},
		{"С параметром -u", strings.Split("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.", "\n"), Options{U: true}, "\nThanks.", require.Equal},
		{"С параметром -i", strings.Split("I LOVE MUSIC.\nI love music.\nI LoVe MuSiC.\n\nI love MuSIC of Kartik.\nI love music of kartik.\nThanks.\nI love music of kartik.\nI love MuSIC of Kartik.", "\n"), Options{I: true}, "I LOVE MUSIC.\n\nI love MuSIC of Kartik.\nThanks.\nI love music of kartik.", require.Equal},
		{"С параметром -f num", strings.Split("We love music.\nI love music.\nThey love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.", "\n"), Options{NumFields: 1}, "We love music.\n\nI love music of Kartik.\nThanks.", require.Equal},
		{"С параметром -s num", strings.Split("I love music.\nA love music.\nC love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.", "\n"), Options{NumChars: 1}, "I love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.", require.Equal},
		{"Нет входных данных", []string{}, Options{}, "", require.Equal},
		{"Неправильные параметры", strings.Split("I love music.\nI love music.", "\n"), Options{C: true, D: true, U: true}, "", require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, Uniq(tt.inputLines, tt.inputOptions))
		})
	}
}
