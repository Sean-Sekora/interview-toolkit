package main

import (
	"coderpad/model"
	"fmt"
)

func FindFood(pond *model.Pond, youAreHere *model.LillyPad) *model.LillyPad {
	// TODO: Implement this function
	return nil
}

func main() {

	var paths = []string{
		"1,2,4",
		"2,1,4",
		"3,6,9,10",
		"4,1,2,5,6",
		"5,4,7",
		"6,3,4,10",
		"7,5,10",
		"8,3,9",
		"9,3,8",
		"10,6,7",
	}

	pond := model.NewPond(paths)

	placeOfBirth := pond.Birth()
	fmt.Printf("Place of birth: %d\n", placeOfBirth.Id)

	foodLocation := FindFood(&pond, placeOfBirth)

	if foodLocation == nil {
		fmt.Println("The frog starved")
	} else {
		fmt.Printf("Found food at: %d\n", foodLocation.Id)
	}
}
