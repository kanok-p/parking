package service

import (
	"fmt"
)

func (impl *implementation) RunCommand(commands []CommandInput) {
	for _, command := range commands {
		result := impl.Stdin(command)
		fmt.Println(result)
	}
}
