package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseStr(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		expect    []string
		assertion require.ComparisonAssertionFunc
	}{
		{"", "1 + 2    + 3   ", []string{"1", "+", "2", "+", "3"}, require.Equal},
		{"", "15 + 200    + 3   ", []string{"15", "+", "200", "+", "3"}, require.Equal},
		{"", "3 + 4 * 2 / (1 - 5)*2", []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "*", "2"}, require.Equal},
		{"", "-3 + 4 * 2 / (1 - 5)*2", []string{"-", "3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "*", "2"}, require.Equal},
		{"", "((1+2*3))*((2-1)/1)", strings.Split("((1+2*3))*((2-1)/1)", ""), require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, parseStr(tt.inputStr))
		})
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		expect    []string
		assertion require.ComparisonAssertionFunc
	}{
		//{"", "1 + 2", []string{"1", "2", "+"}, require.Equal},
		//{"", "1 + 2    + 3   ", []string{"1", "2", "+", "3", "+"}, require.Equal},
		//{"", "15 + 200    + 3   ", []string{"15", "200", "+", "3", "+"}, require.Equal},
		{"", "3 + 4 * 2 / (1 - 5)*2", strings.Split("342*15-/2*+", ""), require.Equal},
		//{"", "((1+2*3))*((2-1)/1)", strings.Split("123*+21-1/*", ""), require.Equal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, convertExpToRPN(tt.inputStr))
		})
	}
}
