package prewire

import (
	"fmt"
	"os"

	"github.com/cmaher/prewire"
)

func main() {
	if err := prewire.PrewireCommand(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
