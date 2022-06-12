package exutils

import (
	"os"
	"strings"
)

// WORKERTYPES
// Encode the type of machine, so we know what GPU and how much vram etc
const (
	P100_16GB_X1 = iota // 1X P100 with 16GB VRAM
)

// Checks for exactly 1 arg from stdin and returns it as string
// Prints error if not exactly 1 argument received
func InitializeSecretFromArgument() string {
	if len(os.Args) < 2 || len(os.Args) > 2 { // Check arguments
		println("Error: You must supply EXACTLY one argument (the GPU_WORER auth token) on startup.")
		os.Exit(1)
	}
	return strings.TrimSpace(os.Args[1])
}
