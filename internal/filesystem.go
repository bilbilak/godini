package internal

import (
	"os"
	"strings"
)

func saveConfig(path string, lines []string) {
	content := strings.TrimSpace(strings.Join(lines, "\n")) + "\n"
	err := os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		FatalError(err)
	}
}
