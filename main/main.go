package main

import (
	"fmt"
	"os"

	"github.com/AlTavares/coderunner/runner"
)

func main() {
	args := os.Args[1:]
	fmt.Println("command:", args)
	output, executionTime, err := runner.Run(args[0], args[1:])
	fmt.Println(output, executionTime, err)
}
