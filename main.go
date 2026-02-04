package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	inputPath  string
	outputPath string
	showHelp   bool
)

func init() {
	flag.StringVar(&inputPath, "i", "", "Path to input csv file (default: stdin)")
	flag.StringVar(&outputPath, "o", "", "Path to output markdown file (default: stdout)")
	flag.BoolVar(&showHelp, "h", false, "Show how to use tool. Can use stdin/stdout, positional args [1] input [2] output, or flags -i/-o")
}

func getSrc() (io.ReadCloser, error) {
	// flag
	if inputPath != "" {
		return os.Open(inputPath)
	}
	// positional args
	if args := flag.Args(); len(args) > 0 {
		return os.Open(args[0])
	}
	// check if stdin is just an empty terminal
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("No input provided. Usage: csv2md [input.csv] [output.md] or cat data.csv | csv2md")
	}
	// stdin
	return io.NopCloser(os.Stdin), nil
}

func getDest() (io.WriteCloser, error) {
	// flag
	if outputPath != "" {
		return os.Create(outputPath)
	}
	// positional args
	if args := flag.Args(); len(args) > 1 {
		return os.Create(args[1])
	}
	// stdout
	return os.Stdout, nil
}

func readInput() ([][]string, error) {
	var r io.ReadCloser

	r, err := getSrc()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	csvReader := csv.NewReader(r)
	return csvReader.ReadAll()
}

func writeOutput(output []string) error {
	w, err := getDest()
	if err != nil {
		return err
	}
	defer w.Close()

	writer := bufio.NewWriter(w)
	outputString := strings.Join(output, "\n")
	writer.WriteString(outputString)
	writer.Flush()
	return nil
}

func getWidths(input [][]string) (widths []int) {
	if len(input) == 0 {
		return nil
	}
	nCols := len(input[0])
	widths = make([]int, nCols)

	for _, row := range input {
		for idx, val := range row {
			if nChars := len(val); nChars > widths[idx] {
				widths[idx] = nChars
			}
		}
	}

	return widths
}

func transformRow(row []string, widths []int) string {
	var formattedEntries []string
	for idx, val := range row {
		padding := strings.Repeat(" ", widths[idx]-len(val))
		formattedEntries = append(formattedEntries, padding+val)
	}
	formattedLine := strings.Join(formattedEntries, " | ")
	return "| " + formattedLine + " |"
}

func csvToMarkdown(input [][]string) (output []string) {
	if len(input) == 0 {
		return nil
	}
	widths := getWidths(input)

	// Header
	header := input[0]
	output = append(output, transformRow(header, widths))
	// Divider
	var divider []string
	for _, width := range widths {
		dividerBlock := strings.Repeat("-", width+2)
		divider = append(divider, dividerBlock)
	}
	formattedDivider := strings.Join(divider, "|")
	output = append(output, "|"+formattedDivider+"|")
	// Rows
	rows := input[1:]
	for _, row := range rows {
		output = append(output, transformRow(row, widths))
	}
	return output
}

func main() {
	// runs init under the hood
	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	input, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		os.Exit(1)
	}
	if len(input) == 0 {
		fmt.Fprintln(os.Stderr, "Error: empty CSV")
		os.Exit(1)
	}
	output := csvToMarkdown(input)
	err = writeOutput(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
}
