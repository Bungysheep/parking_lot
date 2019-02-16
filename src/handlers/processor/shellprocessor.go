package processor

import (
	"bufio"
	"fmt"
	"handlers/commandmanager"
	"os"
	"strings"
	"utils/constant"
)

// ShellProcessor type
type ShellProcessor struct {
	PS1 string
}

// NewShellProcessor - Create a new Shell Processor
func NewShellProcessor() *ShellProcessor {
	return new(ShellProcessor)
}

// Run - Run the shell process
func (sp *ShellProcessor) Run() {
	reader := bufio.NewReader(os.Stdin)
	cm := commandmanager.NewCommandManager()

	sp.displayPrompt()
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, constant.NewLine)
		input = strings.TrimSpace(input)
		if input != constant.EmptyString {
			switch input {
			case constant.ExitAction:
				os.Exit(2)
			default:
				result, err := cm.Execute(input)
				if err != nil {
					fmt.Println("Error:", err.Error())
				} else {
					fmt.Println(result)
				}
			}
		}

		fmt.Print(constant.NewLine)
		sp.displayPrompt()
	}
}

func (sp *ShellProcessor) displayPrompt() {
	fmt.Print(constant.PS1 + " ")
}
