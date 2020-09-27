package logic

import (
	"fmt"
	"strings"
)

//@Summary This will take decison of input and print the output
//@Description if Iuput format is not proper then this will print, input data is not correct
func TakeDecisonBasedOnInput(input string) {
	//fmt.Println("TakeDecisonBasedOnInput")
	data := strings.Split(input, " ")
	type DecisonPoint struct {
		fn func([]string)
	}
	//map function based of input
	var c = map[string]DecisonPoint{}
	// call function based on input
	if _, ok := c[data[0]]; ok {
		c[data[0]].fn(data)
	} else {
		fmt.Println("Input is not proper")
	}

}
