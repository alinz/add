package main

import (
	"context"
	"fmt"
	"os"

	"selfupdate.blockthrough.com"
)

var (
	Version   = ""
	PublicKey = ""
)

const (
	ownerName = "alinz"
	repoName  = "add"
	execName  = "add"
)

func main() {
	runUpdate()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <number> <number> ...\n", os.Args[0])
		os.Exit(0)
	}

	sum := 0

	for _, arg := range os.Args[1:] {
		var num int
		_, err := fmt.Sscanf(arg, "%d", &num)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		sum += num
	}

	fmt.Printf("Sum: %d\n", sum)
}

func runUpdate() {
	ghToken, ok := os.LookupEnv("ADD_GH_TOKEN")
	if !ok {
		fmt.Fprintf(os.Stderr, "Warning: ADD_GH_TOKEN env is not set, selfupdating is disabled\n")
		return
	}

	selfupdate.Auto(
		context.Background(), // Context
		ownerName,            // Owner Name
		repoName,             // Repo Name
		Version,              // Current Version
		execName,             // Executable Name,
		ghToken,              // Github Token
		PublicKey,            // Public Key
	)
}
