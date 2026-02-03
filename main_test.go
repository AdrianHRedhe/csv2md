package main

import (
	"reflect"
	"testing"
)

func TestGetWidths(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []int
	}{
		{
			name:     "basic table",
			input:    [][]string{{"Name", "Age"}, {"Alice", "30"}},
			expected: []int{5, 3},
		},
		{
			name:     "uneven widths",
			input:    [][]string{{"A", "BB"}, {"CCC", "D"}},
			expected: []int{3, 2},
		},
		{
			name:     "empty input",
			input:    [][]string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getWidths(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTransformRow(t *testing.T) {
	row := []string{"Alice", "30"}
	widths := []int{5, 3}
	expected := "| Alice |  30 |"

	result := transformRow(row, widths)
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestCsvToMarkdown(t *testing.T) {
	input := [][]string{
		{"Name", "Age"},
		{"Alice", "30"},
		{"Bob", "25"},
	}

	result := csvToMarkdown(input)

	if len(result) != 4 { // header + divider + 2 rows
		t.Errorf("expected 4 lines, got %d", len(result))
	}

	// Check it has proper markdown structure
	// padded with one more before name due to len of Alice
	if result[0] != "|  Name | Age |" {
		t.Errorf("unexpected header: %s", result[0])
	}
}
