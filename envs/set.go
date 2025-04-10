package envs

import (
	"bufio"
	"os"
	"strings"
)

func Set(filePath, key, value string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	found := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			lines = append(lines, key+"="+value)
			found = true
		} else {
			lines = append(lines, line)
		}
	}

	if !found {
		lines = append(lines, key+"="+value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}
