package challenges

import (
	"os"

	"time"

	"github.com/AlTavares/coderunner/runner"
)

func Collatz() {
	args := os.Args[1:]
	program := args[0]
	args = args[1:]
	runner.Test(program, args, "10000", "6171 262", 0)
	runner.Test(program, args, "100000", "77031 351", 0)
	runner.Test(program, args, "1000000", "837799 525", 1*time.Second)
}
