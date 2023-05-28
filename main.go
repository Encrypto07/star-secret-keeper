package main

import (
	"fmt"
	"os"

	"github.com/Encrypto07/star-secret-keeper/cmd/secret"
)

func main() {
	if err := secret.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
