package main

import (
	"fmt"
	"hungryfrog/model"
)

func FindFood(pond *model.Pond, youAreHere *model.LillyPad, visited map[int]bool) (*model.LillyPad, error) {
	visited[youAreHere.Id] = true

	if pond.LillyPads[youAreHere.Id].IsThereFood {
		return pond.LillyPads[youAreHere.Id], nil
	}

	fmt.Printf("The frog is on lilly pad %d\n", youAreHere.Id)

	for _, path := range pond.LillyPads[youAreHere.Id].Paths {
		if !visited[path.Id] {
			lillyPad, err := FindFood(pond, path, visited)
			if err != nil {
				return nil, err
			}
			if lillyPad != nil {
				return lillyPad, nil
			}
		} else {
			continue
		}
	}

	return nil, nil
}

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

func main() {
	pond := model.NewPond(paths)

	placeOfBirth := pond.Birth()
	fmt.Printf("Place of birth: %d\n", placeOfBirth.Id)

	foodLocation, err := FindFood(&pond, placeOfBirth, map[int]bool{})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else if foodLocation == nil {
		fmt.Println("The frog starved")
	} else {
		fmt.Printf("Found food at: %d\n", foodLocation.Id)
	}
}
