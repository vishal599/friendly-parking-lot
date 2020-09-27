package logic

import (
	"fmt"
	"strings"
)

const (
	ConstPark                                 = "park"
	ConstStatus                               = "status"
	ConstLeave                                = "leave"
	ConstRegistrationNumbersForCarsWithColour = "registration_numbers_for_cars_with_colour"
	ConstSlotNumbersForCarsWithColour         = "slot_numbers_for_cars_with_colour"
	ConstSlotNumberForRegistrationNumber      = "slot_number_for_registration_number"
	ConstCreateParkingLot                     = "create_parking_lot"
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
	var c = map[string]DecisonPoint{
		ConstCreateParkingLot: DecisonPoint{
			setTotalParkingLotNumber,
		},
		ConstPark: DecisonPoint{
			parkCar,
		},
		ConstStatus: DecisonPoint{
			getParkingStatus,
		},
		ConstLeave: DecisonPoint{
			leaveCar,
		},
		ConstRegistrationNumbersForCarsWithColour: DecisonPoint{
			findCarRegistrationDetailBasedOnColor,
		},
		ConstSlotNumbersForCarsWithColour: DecisonPoint{
			getCarSlotDetailsBasedONColor,
		},
		ConstSlotNumberForRegistrationNumber: DecisonPoint{
			getSlotNumberBasedOnRegistrationDetail,
		},
	}
	// call function based on input
	if _, ok := c[data[0]]; ok {
		c[data[0]].fn(data)
	} else {
		fmt.Println("Input is not proper")
	}

}

//@Summary Set total parking slot and insert data in heap for parking the car nearest to entry point
func setTotalParkingLotNumber(data []string) {

}

//@summary Park the car if available
//@ Description If parking is availble then this will pop the top element from heap and assign that slot to that car and insert data in
// other respective map
func parkCar(data []string) {

}

//@Summary Get parking status
func getParkingStatus(data []string) {

}

//@Summary when parking leave
//@Description  if input was proper then remove data from respective map and insert that slot to heap
//If we give invalid input(like car was not park and want to free), then this will return "Sorry there was no car parked in this slot"
func leaveCar(data []string) {

}

//@Summary find car registration detail based on color
func findCarRegistrationDetailBasedOnColor(data []string) {

}

//@Summary get car slot details based on color
//Description if input color car is not parked then this will return error
func getCarSlotDetailsBasedONColor(data []string) {

}

//@Summary Get slot detail based on registration number
func getSlotNumberBasedOnRegistrationDetail(data []string) {

}
