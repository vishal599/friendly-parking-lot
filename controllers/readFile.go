package controllers

import (
	"bufio"
	"fmt"
	"friendly-parking-lot/logic"
	"log"
	"os"
	"strings"
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

//@Summary Read data from command input and take decison
func ReadDataFromCommandPrompt() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if err != nil {
			fmt.Println("Error in getting input")
		}
		logic.TakeDecisonBasedOnInput(text)

	}
}
