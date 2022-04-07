package utils

import (
	"os"
)

var dataStorageIsDB bool = false

func init() {
	args := os.Args
	errStr := "\nSpecify the data storage location\n-l: locally\n-p: postgresql"
	errStrExample := "\nExample: go run main.go -l"

	if len(args) != 2 {
		panic(errStr + errStrExample)
	}

	saveingMode := args[1]

	switch saveingMode {
	case "-l":
		dataStorageIsDB = false
	case "-p":
		dataStorageIsDB = true
	default:
		panic("\nUnknow parameter" + errStr + errStrExample)
	}
}

func GetDataStorageIsDB() bool {
	return dataStorageIsDB
}
