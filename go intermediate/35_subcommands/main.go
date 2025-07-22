package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	// Define flags for firstSub
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Number of bytes")

	// Define flags for secondSub
	flagSc2 := subCommand2.String("language", "GO", "Enter the language")

	// Ensure at least one argument (subcommand) is provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'firstSub' or 'secondSub' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("Subcommand: firstSub")
		fmt.Println("  Processing:", *firstFlag)
		fmt.Println("  Bytes:", *secondFlag)

	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("Subcommand: secondSub")
		fmt.Println("  Language:", *flagSc2)

	default:
		fmt.Println("Unknown subcommand. Expected 'firstSub' or 'secondSub'")
		os.Exit(1)
	}
}
