package main

import (
	"bass-backend/cli"

	_ "modernc.org/sqlite"
)

func main() {
	cli.ProcessCommandLineArguments()
}
