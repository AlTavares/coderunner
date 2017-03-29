package runner

import (
	"fmt"
	"log"

	"os/exec"
	"strings"
	"time"
)

var Icons = struct {
	success string
	failure string
	info    string
}{
	" ✅ ",
	" 🔴 ",
	" ℹ️  ",
}

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

func Test(program string, args []string, input string, expectedOutput string, expectedExecutionTime time.Duration) (success bool) {
	args = append(args, input)
	output, duration, err := Run(program, args)
	if err != nil {
		log.Fatal(err)
	}
	var icon = Icons.failure

	if output == expectedOutput {
		icon = Icons.success
		success = true
	}
	fmt.Println(icon, "Expected "+expectedOutput+", got "+output)

	if expectedExecutionTime > 0 {
		icon = Icons.success
		if duration > expectedExecutionTime {
			icon = Icons.failure
		}
		fmt.Println(icon, "Executed in", duration, "the maximum expected was", expectedExecutionTime)
		return
	}

	fmt.Println(Icons.info, "Executed in", duration)
	return
}
