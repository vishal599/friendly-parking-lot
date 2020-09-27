package logic

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"

	"fmt"
)

var totalParkingLot int //total no of parking lot

var sloTCarDetailMapping map[string]carDetail // this will store to slot to car detail mapping

var CarRegistrationDetailMapping map[string]slotDetail // this will store registration number to car detail mapping

var CarColorToOtherDetailMapping map[string]map[string]string // thi car detail of same color

type slotDetail struct {
	Slot  string
	Color string
}

type carDeailByColor struct {
	Slot               string
	RegistrationNumber string
}

type carDetail struct {
	RegistrationNumber string
	Color              string
}

type CarFullDetail struct {
	RegistrationNumber string
	Color              string
	Slot               string
}

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
	totalParkingLot, _ = strconv.Atoi(data[1])
	InsertAllSlotWhileCreateingParkingLot(totalParkingLot)
	//push data to heap to calculate closest parking slot while assigning
	fmt.Printf("Created a parking lot with %v slots\n", data[1])
	sloTCarDetailMapping = make(map[string]carDetail, 0)
	CarRegistrationDetailMapping = make(map[string]slotDetail, 0)

	CarColorToOtherDetailMapping = make(map[string]map[string]string, 0)
	//fmt.Println("setTotalParkingLotNumber", len(myNums), totalParkingLot)
}

//@summary Park the car if available
//@ Description If parking is availble then this will pop the top element from heap and assign that slot to that car and insert data in
// other respective map
func parkCar(data []string) {

	//first check available slots should be greater than zero
	if len(myNums) == 0 {
		fmt.Println("Sorry, parking lot is full")
		return
	}
	//if not full then pop nearest slot from heap and assign
	popValue := heap.Pop(&myNums).(MyNum)
	nearestAvailableSlot := popValue.Val
	slotNum := strconv.Itoa(nearestAvailableSlot)
	//insert data in respective map
	SlotToREgistrationMapping(data, slotNum)
	RegistrationToSlotMapping(data, slotNum)
	ColorToCarDetailMapping(data, slotNum)
	fmt.Println("Allocated slot number ", slotNum)
	//fmt.Println("parkCar", len(myNums))
}

// @Summary Slot To Car Detail Maaping
func SlotToREgistrationMapping(data []string, slot string) {
	input := carDetail{}
	input.Color = data[2]
	input.RegistrationNumber = data[1]
	//store detail in format such that we can find which car if park in that slot
	sloTCarDetailMapping[slot] = input
}

//@Summary Registration to car slot Car detail mapping
func RegistrationToSlotMapping(data []string, slot string) {
	input := slotDetail{}
	input.Slot = slot
	input.Color = data[2]
	CarRegistrationDetailMapping[data[1]] = input
}

//@Summary Color to detail mapping of that color
func ColorToCarDetailMapping(data []string, slot string) {
	input := make(map[string]string)
	input[data[1]] = slot

	//check that color car already parked or not
	val, ok := CarColorToOtherDetailMapping[data[2]]
	if ok {
		val[data[1]] = slot
		CarColorToOtherDetailMapping[data[2]] = val
	} else {
		CarColorToOtherDetailMapping[data[2]] = input
	}
}

//@Summary Get parking status
func getParkingStatus(data []string) {
	//CarFullDetail
	//
	output := make([]CarFullDetail, 0)
	carParkingDetails := CarFullDetail{}
	//read data from map
	for key, val := range sloTCarDetailMapping {
		carParkingDetails.Slot = key
		carParkingDetails.RegistrationNumber = val.RegistrationNumber
		carParkingDetails.Color = val.Color
		output = append(output, carParkingDetails)
	}

	//sort output by slots
	sort.SliceStable(output, func(i, j int) bool {
		val1, _ := strconv.Atoi(output[i].Slot)
		val2, _ := strconv.Atoi(output[j].Slot)
		return val1 < val2
	})
	fmt.Println("Slot No.    Registration No        Colour")
	for _, val := range output {
		fmt.Printf("%v           %v            %v\n", val.Slot, val.RegistrationNumber, val.Color)
	}
}

//@Summary when parking leave
//@Description  if input was proper then remove data from respective map and insert that slot to heap
//If we give invalid input(like car was not park and want to free), then this will return "Sorry there was no car parked in this slot"
func leaveCar(data []string) {
	//remove from all map and insert in to heap for next allocation

	//chekc this slot was already filled or not
	carData, ok := sloTCarDetailMapping[data[1]]
	// if this was already free then return error
	if !ok {
		fmt.Println("Sorry there was no car parked in this slot")
		return
	}
	delete(sloTCarDetailMapping, data[1])
	//remove from resitration number to other  detail mapping
	delete(CarRegistrationDetailMapping, carData.RegistrationNumber)

	//remove from color to car detail mapping
	removeCarFromColorMap(carData.RegistrationNumber, carData.Color)

	//push to heap for next avilabilirt
	slotValToInt, _ := strconv.Atoi(data[1])
	availableSlot := MyNum{Val: slotValToInt, Count: totalParkingLot + 1 - slotValToInt}
	// move this slot to heap for availability
	heap.Push(&myNums, availableSlot)
	fmt.Printf("Slot number %v is free\n", data[1])
	//fmt.Println("leaveCar", len(myNums))  // available slot
}

//@Summary remove car from color to car detail mapping while free the car
func removeCarFromColorMap(color, registrationNo string) {
	regisTrationSlotMapping := CarColorToOtherDetailMapping[color]
	delete(regisTrationSlotMapping, registrationNo)
	CarColorToOtherDetailMapping[color] = regisTrationSlotMapping
}

//@Summary find car registration detail based on color
func findCarRegistrationDetailBasedOnColor(data []string) {
	parkedCarOfSpecificColor := CarColorToOtherDetailMapping[data[1]]
	registrationNum := make([]string, 0)
	if len(parkedCarOfSpecificColor) == 0 {
		fmt.Println("Sorry No Car parked of this color")
	} else {
		for key, _ := range parkedCarOfSpecificColor {
			registrationNum = append(registrationNum, key)
		}
	}
	allRegistrationNum := strings.Join(registrationNum, ",")
	fmt.Println(allRegistrationNum)
}

//@Summary get car slot details based on color
//Description if input color car is not parked then this will return error
func getCarSlotDetailsBasedONColor(data []string) {
	parkedCarOfSpecificColor := CarColorToOtherDetailMapping[data[1]]
	allSlots := make([]string, 0)
	// if that color car is not parked
	if len(parkedCarOfSpecificColor) == 0 {
		fmt.Println("Sorry No Car parked of this color")
	} else {
		for _, val := range parkedCarOfSpecificColor {
			allSlots = append(allSlots, val)
		}
	}
	slots := strings.Join(allSlots, ",")
	fmt.Println(slots)
}

//@Summary Get slot detail based on registration number
func getSlotNumberBasedOnRegistrationDetail(data []string) {
	registrationSlotData, ok := CarRegistrationDetailMapping[data[1]]
	if !ok {
		fmt.Println("Not found")
	} else {
		fmt.Println(registrationSlotData.Slot)

	}
}
