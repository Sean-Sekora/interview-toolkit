package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// These are for manual testing. Set to 0 to disable
var foodLocationOverride = 0
var birthPlaceOverride = 0

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
			pond.LillyPads[l].Paths = append(pond.LillyPads[l].Paths, pond.LillyPads[n])
		}
	}

	// Place food
	foodLocation := pond.foodLocation()
	fmt.Printf("Food location: %d\n", foodLocation.Id)
	foodLocation.IsThereFood = true

	return pond
}

func (pond *Pond) Birth() *LillyPad {
	for {
		lillyPad := pond.birthPlace()

		if !lillyPad.IsThereFood {
			return lillyPad
		}
	}
}

func (pond *Pond) randomLillyPad() *LillyPad {
	return pond.LillyPads[rand.Intn(len(pond.LillyPads)-1)+1]
}

func (pond *Pond) foodLocation() *LillyPad {
	if foodLocationOverride > 0 {
		return pond.LillyPads[foodLocationOverride]
	} else {
		return pond.randomLillyPad()
	}
}

func (pond *Pond) birthPlace() *LillyPad {
	if birthPlaceOverride > 0 {
		return pond.LillyPads[birthPlaceOverride]
	} else {
		return pond.randomLillyPad()
	}
}
