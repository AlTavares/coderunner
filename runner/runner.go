package runner

import (
	"fmt"
	"log"

	"os/exec"
	"strings"
	"time"
)

func Run(program string, args []string) (string, time.Duration, error) {
	fmt.Println("\nCommand:", program, args)
	start := time.Now()
	out, err := exec.Command(program, args...).Output()
	timeElapsed := time.Since(start)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	return strings.TrimSpace(fmt.Sprintf("%s", out)), timeElapsed, err
}

func Test(program string, args []string, input string, expectedOutput string, expectedExecutionTime time.Duration) {
	args = append(args, input)
	output, duration, err := Run(program, args)
	if err != nil {
		log.Fatal(err)
	}
	var icon = " ✅ "

	if output != expectedOutput {
		icon = " 🔴 "
	}
	fmt.Println(icon, "Expected "+expectedOutput+", got "+output)

	if expectedExecutionTime > 0 {
		icon = " ✅ "
		if duration > expectedExecutionTime {
			icon = " 🔴 "
		}
		fmt.Println(icon, "Executed in", duration, "the maximum expected was", expectedExecutionTime)
		return
	}

	fmt.Println(" ℹ️  Executed in", duration)

}
