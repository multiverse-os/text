package text

import (
	"fmt"
)

func Box(message string) (output string) {
	messageLength := len(message)
	output += fmt.Sprintf("╭")
	for i := 0; i < l+2; i++ {
		output += fmt.Sprintf("─")
	}
	output += fmt.Sprintf("╮\n")
	output += fmt.Sprintf("│ %v │\n", message)
	output += fmt.Sprintf("╰")
	for i := 0; i < l+2; i++ {
		output += fmt.Sprintf("─")
	}
	output += fmt.Sprintf("╯\n")
  return  output
}
