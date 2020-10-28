package service

import (
	"os"
	"strings"
)

func (impl *implementation) OpenFile(path string) (commands []CommandInput, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []string
	buf := make([]byte, 1000)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			return nil, err
		}

		lines = strings.Split(string(buf[:n]), "\n")
	}
	commands = splitLine(lines)
	return commands, nil
}
func splitLine(lines []string) (Commands []CommandInput) {
	for _, line := range lines {
		comm := strings.Split(line, " ")
		Commands = append(Commands, splitCommand(comm))
	}

	return Commands
}

func splitCommand(commands []string) (command CommandInput) {
	command = CommandInput{LengthOfParam: len(commands)}
	for i, col := range commands {
		switch i {
		case 0:
			command.Action = col
		case 1:
			command.StParam = col
		case 2:
			command.NdParam = col
		}
	}

	return command
}
