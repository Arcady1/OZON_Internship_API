package utils

import (
	"log"
	"os"
)

var dataStorageIsDB bool = false

func init() {
	args := os.Args
	errStr := "Specify the data storage location with the last parameter\nl: locally\np: postgresql"
	errStrExample := "\n\nExample 1: go run main.go l\nExample 2: go test -v -args l\n\n"
	mess := "The data is now saved "

	saveingMode := args[len(args)-1]

	switch saveingMode {
	case "l":
		dataStorageIsDB = false
		log.Println(mess + "LOCALLY")
	case "p":
		dataStorageIsDB = true
		log.Println(mess + "IN DATABASE")
	default:
		dataStorageIsDB = false
		log.Println(errStr + errStrExample + mess + "LOCALLY")
	}
}

func GetDataStorageIsDB() bool {
	return dataStorageIsDB
}

func SetDataStorageIsDB(status bool) {
	dataStorageIsDB = status
}
