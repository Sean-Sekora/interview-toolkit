package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Pond struct {
	LillyPads map[int]*LillyPad
}

func NewPond(paths []string) Pond {
	pond := Pond{
		map[int]*LillyPad{},
	}

	// Create the lilly pads
	for _, path := range paths {
		s := strings.Split(path, ",")
		i, _ := strconv.Atoi(s[0])
		pond.LillyPads[i] = &LillyPad{Id: i}
	}

	// Add neighbors
	for _, path := range paths {
		s := strings.Split(path, ",")
		l, _ := strconv.Atoi(s[0])
		for i := 1; i < len(s); i++ {
			n, _ := strconv.Atoi(s[i])
			pond.LillyPads[l].Neighbors = append(pond.LillyPads[l].Neighbors, pond.LillyPads[n])
		}
	}

	// Place food
	foodLocation := pond.randomLillyPad()
	fmt.Printf("Food location: %d\n", foodLocation.Id)
	foodLocation.IsThereFood = true

	return pond
}

func (pond *Pond) Birth() *LillyPad {
	for {
		lillyPad := pond.randomLillyPad()

		if !lillyPad.IsThereFood {
			return lillyPad
		}
	}
}

func (pond *Pond) randomLillyPad() *LillyPad {
	return pond.LillyPads[rand.Intn(len(pond.LillyPads)-1)+1]
}
