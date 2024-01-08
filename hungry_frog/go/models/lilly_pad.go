package models

type LillyPad struct {
	Id          int
	IsThereFood bool
	Neighbors   []*LillyPad
}
