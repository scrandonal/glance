package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/glanceapp/glance/internal/glance"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
)

func main() {
	// Default config path changed to match my personal setup convention
	configPath := flag.String("config", "config/glance.yml", "Path to the configuration file")
	showVersion := flag.Bool("version", false, "Print version information and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("glance %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	app, err := glance.New(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
