package main

import (
	"fmt"
	"hungryfrog/models"
)

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
	pond := models.NewPond(paths)
	placeOfBirth := pond.Birth()
	fmt.Printf("Place of birth: %d\n", placeOfBirth.Id)
	foodLocation := DepthFirstSearch(&pond, placeOfBirth.Id, map[int]bool{})
	fmt.Printf("Found food at: %d\n", foodLocation.Id)
}

func DepthFirstSearch(pond *models.Pond, vertix int, visited map[int]bool) *models.LillyPad {
	visited[vertix] = true
	fmt.Printf("Hop to %d\n", vertix)

	if pond.LillyPads[vertix].IsThereFood {
		return pond.LillyPads[vertix]
	}

	for _, neighbor := range pond.LillyPads[vertix].Neighbors {
		var lillyPad *models.LillyPad
		if !visited[neighbor.Id] {
			lillyPad = DepthFirstSearch(pond, neighbor.Id, visited)
			if lillyPad != nil {
				return lillyPad
			} else {
				fmt.Printf("Hop to %d\n", vertix)
			}
		} else {
			continue
		}
	}

	return nil
}
