package main

import (
	"flag"
	"os"
	"triforce-blitz/routes"
)

func LoadConfiguration() {
	parseConfigurationFile()
	parseEnvironmentVariables()
	parseCommandLineFlags()
}

func parseConfigurationFile() {
	// TODO: read from a configuration file
}

func parseEnvironmentVariables() {
	routes.SeedsDir = os.Getenv("TRIFORCE_BLITZ_SEEDS_DIR")
}

func parseCommandLineFlags() {
	flag.Parse()
}

func init() {
	flag.StringVar(&routes.SeedsDir, "seeds-dir", "", "directory in which to store generated seeds")
	flag.StringVar(&routes.SeedsDir, "S", "", "directory in which to store generated seeds (shorthand)")
}
