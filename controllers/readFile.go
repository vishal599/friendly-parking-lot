package controllers

import (
	"bufio"
	"friendly-parking-lot/logic"
	"log"
	"os"
)

//@Summary Read from input txt file and take decison what to do on which input
func ReadDataFromFileAndTakeDecison(sourceFile string) {
	body, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer body.Close()

	scanner := bufio.NewScanner(body)
	//line by line input
	for scanner.Scan() {
		//take decison based on data
		logic.TakeDecisonBasedOnInput(scanner.Text())
	}
	return
}
