package main

import (
	"friendly-parking-lot/controllers"
	"os"
)

func main() {
	//if args length is >1  means input is file
	if len(os.Args) > 1 {
		sourceFileName := os.Args[1]
		//fmt.Println("read data from file", sourceFileName)
		controllers.ReadDataFromFileAndTakeDecison(sourceFileName)
		return

	}
	//if input given from command
	controllers.ReadDataFromCommandPrompt()
	return
}
