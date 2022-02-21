package cli

import (
	"flag"
	"fmt"
	"os"

	"go_gin_react_blog/internal/logging"
)

func usage() {
	fmt.Print(`This program runs go_gin_react_blog backend server.
Usage:
go_gin_react_blog [arguments]
Supported arguments:
`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() string {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
	return *env
}
