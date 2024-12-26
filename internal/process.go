package internal

import (
	"fmt"
	"regexp"
	"strings"
)

type setting struct {
	key     string
	value   string
	section string
}

type operation int

const (
	GET operation = iota
	SET
	UNSET
	DELETE
)

func Process(operation operation, args []string) {
	config, settings := fetchInput(args)

	for _, s := range settings {
		setting := parseSetting(s)

		if setting == nil {
			continue
		}

		foundSection := false
		foundKey := false
		lastSection := ""

		if len(Section) > 0 {
			setting.section = Section
		}

		for n, line := range config {
			if strings.Contains(line, "[") && strings.Contains(line, "]") {
				sectionEnd := strings.Index(line, "]")
				lastSection = line[1:sectionEnd]
			}

			if len(setting.section) > 0 {
				if !foundSection {
					if strings.HasPrefix(line, "["+setting.section+"]") {
						foundSection = true
					}

					continue
				} else if strings.HasPrefix(line, "[") {
					if operation == SET && !foundKey {
						config = insertLine(config, n, setting.key+"="+setting.value)
						foundKey = true
					}

					foundSection = false
					break
				}
			} else if !All && strings.HasPrefix(line, "[") {
				break
			}

			re := regexp.MustCompile(`^;?\s*` + regexp.QuoteMeta(setting.key) + `\s*=`)

			if re.MatchString(line) {
				switch operation {
				case GET:
					if Full {
						prefix := ""

						if len(lastSection) > 0 {
							prefix = "[" + lastSection + "]"
						}

						fmt.Println(strings.Replace(line, setting.key, prefix+setting.key, 1))
					} else {
						fmt.Println(line)
					}
				case SET: // TODO: "=" means new value although empty
					foundKey = true
					line = strings.TrimPrefix(line, ";")

					if len(setting.value) > 0 {
						parts := strings.SplitN(line, "=", 2)

						if len(parts) > 1 {
							line = strings.TrimSpace(parts[0]) + "=" + setting.value
						}
					}

					config[n] = line
				case UNSET:
					line = strings.TrimPrefix(line, ";")
					config[n] = ";" + line
				case DELETE:
					config = deleteLine(config, n)
				}
			}
		}

		if operation == SET && !foundKey {
			entry := setting.key + "=" + setting.value

			if len(setting.section) > 0 {
				if foundSection {
					config = append(config, "", entry)
				} else {
					config = append(config, "", "["+setting.section+"]", entry)
				}
			} else {
				config = append([]string{entry, ""}, config...)
			}
		}

		saveConfig(args[0], config)
	}
}

func parseSetting(input string) *setting {
	re := regexp.MustCompile(`^(?:\[(?P<section>[a-zA-Z0-9._-]+)\])?(?P<key>[a-zA-Z0-9_-]+)(?:=(?P<value>.*))?$`)
	matches := re.FindStringSubmatch(input)

	if matches == nil {
		return nil
	}

	setting := setting{}

	for i, name := range re.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}

		switch name {
		case "section":
			setting.section = matches[i]
		case "key":
			setting.key = matches[i]
		case "value":
			setting.value = matches[i]
		}
	}

	return &setting
}

func insertLine(lines []string, index int, value string) []string {
	if index < 0 || index > len(lines) {
		return lines
	}

	lines = append(lines[:index+2], lines[index:]...)
	lines[index] = value
	lines[index+1] = ""

	return lines
}

func deleteLine(lines []string, index int) []string {
	if index < 0 || index >= len(lines) {
		return lines
	}

	return append(lines[:index], lines[index+1:]...)
}
