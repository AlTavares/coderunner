package runner

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func Run(program string, args []string) (string, time.Duration, error) {
	log.Println("Command:", program, args)
	start := time.Now()
	out, err := exec.Command(args[0], args[1:]...).Output()
	timeElapsed := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(timeElapsed)
	return fmt.Sprintf("%s", out), timeElapsed, err
}
