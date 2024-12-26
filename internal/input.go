package internal

import (
	"bufio"
	"os"
	"strings"
)

func fetchInput(args []string) ([]string, []string) {
	config := loadConfig(args[0])

	var settings []string

	if len(args) > 1 {
		settings = args[1:]
	}

	input := readStdIn()

	if len(input) > 0 {
		settings = append(settings, input...)
	}

	return config, settings
}

func loadConfig(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		FatalError(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		FatalError(err)
	}

	return lines
}

func readStdIn() []string {
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return []string{}
	}

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		FatalError(err)
	}

	return lines
}
